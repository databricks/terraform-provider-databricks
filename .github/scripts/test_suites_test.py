import tempfile
from pathlib import Path
import unittest

from test_suites import replay_matrix, replay_suites, unit_packages


class TestSuiteDiscovery(unittest.TestCase):
    def setUp(self):
        self.directory = tempfile.TemporaryDirectory()
        self.addCleanup(self.directory.cleanup)
        self.root = Path(self.directory.name)

    def create_file(self, path):
        file = self.root / path
        file.parent.mkdir(parents=True, exist_ok=True)
        file.touch()

    def test_discovers_nested_suites_once_in_stable_order(self):
        self.create_file("qa/replay/resources/group/replay_test.go")
        self.create_file("qa/replay/resources/group/extra_test.go")
        self.create_file("qa/replay/features/provider_config/replay_test.go")
        self.create_file("qa/replay/clienttest/client.go")
        self.assertEqual(replay_suites(self.root), [
            "./qa/replay/features/provider_config",
            "./qa/replay/resources/group",
        ])

    def test_ignores_go_fixture_and_hidden_directories_and_files(self):
        self.create_file("qa/replay/resources/group/replay_test.go")
        for name in ["testdata", "vendor", ".hidden", "_ignored"]:
            self.create_file(f"qa/replay/{name}/nested/replay_test.go")
        self.create_file("qa/replay/ignored/.hidden_test.go")
        self.create_file("qa/replay/ignored/_ignored_test.go")
        self.assertEqual(replay_suites(self.root), ["./qa/replay/resources/group"])

    def test_new_suite_is_discovered_automatically(self):
        self.create_file("qa/replay/resources/group/replay_test.go")
        self.assertEqual(len(replay_suites(self.root)), 1)
        self.create_file("qa/replay/resources/job/replay_test.go")
        self.assertEqual(replay_suites(self.root), [
            "./qa/replay/resources/group",
            "./qa/replay/resources/job",
        ])

    def test_rejects_empty_replay_matrix(self):
        with self.assertRaisesRegex(ValueError, "No replay test suites"):
            replay_suites(self.root)

    def test_matrix_groups_all_suites_once_into_eight_nonempty_jobs(self):
        for index in range(257):
            self.create_file(f"qa/replay/resources/suite{index}/replay_test.go")
        groups = replay_matrix(self.root)
        self.assertEqual([group["shard"] for group in groups], list(range(1, 9)))
        self.assertTrue(all(group["packages"] for group in groups))
        packages = [package for group in groups for package in group["packages"]]
        self.assertEqual(sorted(packages), replay_suites(self.root))

    def test_matrix_balances_yaml_scenarios_before_assigning_small_suites(self):
        for suite, count in [("a", 8), ("b", 6), ("c", 4), ("d", 2)]:
            self.create_file(f"qa/replay/{suite}/replay_test.go")
            for index in range(count):
                extension = "yaml" if index % 2 else "yml"
                self.create_file(f"qa/replay/{suite}/testdata/{index}.{extension}")
        self.assertEqual(replay_matrix(self.root, shard_count=2), [
            {"shard": 1, "packages": ["./qa/replay/a", "./qa/replay/d"]},
            {"shard": 2, "packages": ["./qa/replay/b", "./qa/replay/c"]},
        ])

    def test_matrix_does_not_create_empty_jobs_for_small_repositories(self):
        self.create_file("qa/replay/resources/group/replay_test.go")
        self.assertEqual(replay_matrix(self.root), [
            {"shard": 1, "packages": ["./qa/replay/resources/group"]},
        ])

    def test_unit_packages_exclude_suites_but_keep_replay_helpers(self):
        self.create_file("qa/replay/resources/group/replay_test.go")
        directories = ["catalog", "qa/replay/resources/group", "qa/replay/clienttest", "qa/terraformtest"]
        packages = [f"example.com/provider/{path} {self.root / path}\n" for path in directories]
        self.assertEqual(unit_packages(self.root, packages), [
            "example.com/provider/catalog",
            "example.com/provider/qa/replay/clienttest",
            "example.com/provider/qa/terraformtest",
        ])

    def test_rejects_empty_unit_selection(self):
        self.create_file("qa/replay/resources/group/replay_test.go")
        with self.assertRaisesRegex(ValueError, "No non-replay Go packages"):
            unit_packages(self.root, [])


if __name__ == "__main__":
    unittest.main()
