"""Regression tests for CI package selection, without requiring Go or Terraform."""

import subprocess
import sys
import tempfile
import unittest
from contextlib import redirect_stderr, redirect_stdout
from io import StringIO
from pathlib import Path
from unittest.mock import patch

import select_test_packages as selector


class SelectTestPackagesTest(unittest.TestCase):
    def setUp(self):
        directory = tempfile.TemporaryDirectory(prefix="test selection ")
        self.addCleanup(directory.cleanup)
        self.directory = Path(directory.name)

    def package(self, name, has_tests="test", source_size=1, fixtures=0):
        directory = self.directory / name
        directory.mkdir(parents=True)
        (directory / "main.go").write_bytes(b"x" * source_size)
        if fixtures:
            testdata = directory / "testdata"
            testdata.mkdir()
            for index in range(fixtures):
                (testdata / f"{index}.yaml").write_text("")
        return [f"example.org/provider/{name}", has_tests, str(directory)]

    def groups(self, packages, suite, count=2):
        with redirect_stderr(StringIO()):
            return [
                selector.select_packages(packages, suite, shard, count)
                for shard in range(1, count + 1)
            ]

    def run_cli(self, arguments, packages=(), error=None):
        inventory = "".join(" ".join(row) + "\n" for row in packages)
        stdout, stderr = StringIO(), StringIO()
        # Go metadata is an external boundary; selection and parsing run unchanged.
        with (
            patch.object(sys, "argv", ["select_test_packages.py", *arguments]),
            patch.object(
                selector.subprocess, "check_output",
                return_value=inventory, side_effect=error,
            ),
            redirect_stdout(stdout), redirect_stderr(stderr),
        ):
            try:
                selector.main()
                code = 0
            except SystemExit as exit_status:
                code = exit_status.code
        return code, stdout.getvalue(), stderr.getvalue()

    def test_every_package_is_assigned_once_with_coverage_helpers_retained(self):
        unit = self.package("unit")
        generated = self.package("generated", has_tests="none")
        replay = self.package("qa/replay/lifecycle", fixtures=3)
        client = self.package("qa/replay/client", fixtures=0)
        helper = self.package("qa/replay/helper", has_tests="none")
        packages = [unit, generated, replay, client, helper]

        units = sum(self.groups(packages, "unit"), [])
        replays = sum(self.groups(packages, "replay"), [])

        self.assertCountEqual(units, [unit[0], generated[0], helper[0]])
        self.assertCountEqual(replays, [replay[0], client[0]])
        self.assertCountEqual(units + replays, [row[0] for row in packages])

    def test_assignment_is_independent_of_discovery_order_and_weight_ties(self):
        packages = [self.package(f"unit{index}") for index in range(4)]
        packages += [self.package(f"qa/replay/suite{index}") for index in range(4)]
        for suite in ("unit", "replay"):
            with self.subTest(suite=suite):
                self.assertEqual(
                    self.groups(packages, suite),
                    self.groups(list(reversed(packages)), suite),
                )

    def test_heavy_packages_are_balanced_for_both_suites(self):
        for suite in ("unit", "replay"):
            prefix = f"{suite}/" if suite == "unit" else "qa/replay/"
            packages = [
                self.package(prefix + str(weight), source_size=weight, fixtures=weight)
                for weight in (8, 4, 3, 1)
            ]
            with self.subTest(suite=suite):
                groups = self.groups(packages, suite)
                self.assertEqual(groups[0], [packages[0][0]])
                self.assertCountEqual(groups[1], [row[0] for row in packages[1:]])

    def test_unit_weight_counts_direct_go_sources(self):
        package = self.package("unit", source_size=10)
        directory = Path(package[2])
        (directory / "main_test.go").write_bytes(b"x" * 5)
        (directory / "README.md").write_bytes(b"x" * 100)
        (directory / "nested").mkdir()
        (directory / "nested/ignored.go").write_bytes(b"x" * 100)
        self.assertEqual(selector.unit_weight(directory), 15)

    def test_replay_weight_counts_only_direct_yaml_files(self):
        package = self.package("qa/replay/fixtures", fixtures=1)
        directory = Path(package[2])
        testdata = directory / "testdata"
        (testdata / "second.yml").write_text("")
        (testdata / "README.md").write_text("")
        (testdata / "directory.yaml").mkdir()
        (testdata / "nested").mkdir()
        (testdata / "nested/ignored.yaml").write_text("")
        self.assertEqual(selector.replay_weight(directory), 2)

    def test_missing_or_empty_sources_and_fixtures_have_minimum_weight(self):
        empty = self.directory / "empty"
        empty.mkdir()
        for directory in (empty, self.directory / "missing"):
            with self.subTest(directory=directory):
                self.assertEqual(selector.unit_weight(directory), 1)
                self.assertEqual(selector.replay_weight(directory), 1)

    def test_cli_parses_directories_with_spaces_and_prints_only_packages(self):
        package = self.package("unit")
        code, stdout, stderr = self.run_cli(
            ["unit", "--shard", "1", "--shard-count", "1"], [package]
        )
        self.assertEqual(code, 0)
        self.assertEqual(stdout, package[0] + "\n")
        self.assertIn("Unit shard 1", stderr)

    def test_cli_rejects_invalid_shards_and_suites(self):
        for suite in ("unit", "replay"):
            for shard, count in (
                ("0", "2"), ("3", "2"), ("1", "0"), ("1", "-1"), ("all", "2"),
            ):
                with self.subTest(suite=suite, shard=shard, count=count):
                    code, stdout, _ = self.run_cli(
                        [suite, "--shard", shard, "--shard-count", count]
                    )
                    self.assertEqual(code, 2)
                    self.assertEqual(stdout, "")
        for arguments in (["unknown", "--shard-count", "2"], ["unit"]):
            with self.subTest(arguments=arguments):
                self.assertEqual(self.run_cli(arguments)[0], 2)

    def test_cli_rejects_empty_selections_including_unpopulated_shards(self):
        unit = self.package("unit")
        for suite, packages, shard in (
            ("unit", [], "1"),
            ("replay", [], "1"),
            ("unit", [unit], "2"),
            ("replay", [unit], "1"),
        ):
            with self.subTest(suite=suite, packages=packages, shard=shard):
                code, stdout, stderr = self.run_cli(
                    [suite, "--shard", shard, "--shard-count", "2"], packages
                )
                self.assertEqual(code, 1)
                self.assertEqual(stdout, "")
                self.assertIn("No packages selected", stderr)

    def test_go_discovery_failures_propagate(self):
        with self.assertRaises(subprocess.CalledProcessError):
            self.run_cli(
                ["unit", "--shard-count", "1"],
                error=subprocess.CalledProcessError(1, ["go", "list"]),
            )


if __name__ == "__main__":
    unittest.main()
