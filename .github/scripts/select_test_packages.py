#!/usr/bin/env python3
"""Print the Go packages assigned to one job in the fixed test matrix."""

import argparse
import subprocess
import sys
from pathlib import Path


def load_packages():
    # Let Go handle build constraints and ignored directories when discovering tests.
    template = (
        "{{.ImportPath}} "
        "{{if or .TestGoFiles .XTestGoFiles}}test{{else}}none{{end}} {{.Dir}}"
    )
    output = subprocess.check_output(["go", "list", "-f", template, "./..."], text=True)
    return [line.split(maxsplit=2) for line in output.splitlines()]


def replay_weight(directory):
    # Direct YAML fixtures approximate suite runtime without maintaining timing data.
    fixtures = (Path(directory) / "testdata").glob("*")
    count = sum(
        path.is_file() and path.suffix in (".yaml", ".yml") for path in fixtures
    )
    # Client replay tests may have no lifecycle fixtures but still need a shard.
    return max(1, count)


def select_packages(packages, suite, shard, shard_count):
    unit_packages = []
    replay_packages = []
    for import_path, has_tests, directory in packages:
        if "/qa/replay/" in import_path and has_tests == "test":
            replay_packages.append((import_path, directory))
        else:
            # Keep replay helper packages without tests in the normal coverage report.
            unit_packages.append(import_path)

    if suite == "unit":
        return unit_packages

    weighted = [
        (replay_weight(directory), name) for name, directory in replay_packages
    ]
    groups = [[] for _ in range(shard_count)]
    loads = [0] * shard_count
    # Largest suites go first to avoid concentrating expensive packages in one job.
    # Name and shard-index tie breaks keep independent jobs' assignments identical.
    for weight, name in sorted(weighted, key=lambda item: (-item[0], item[1])):
        target = min(range(shard_count), key=lambda index: loads[index])
        groups[target].append(name)
        loads[target] += weight

    selected = groups[shard - 1]
    print(
        f"Replay shard {shard}: {len(selected)} packages, weight {loads[shard - 1]}",
        file=sys.stderr,
    )
    return selected


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("suite", choices=("unit", "replay"))
    parser.add_argument(
        "--shard", default="1", help="1-based replay shard; ignored for unit"
    )
    parser.add_argument("--shard-count", type=int, default=8)
    args = parser.parse_args()
    if args.shard_count < 1:
        parser.error("shard count must be positive")
    shard = 1
    if args.suite == "replay":
        try:
            shard = int(args.shard)
        except ValueError:
            parser.error("replay shard must be an integer")
        if not 1 <= shard <= args.shard_count:
            parser.error("replay shard must be between 1 and shard count")

    selected = select_packages(load_packages(), args.suite, shard, args.shard_count)
    # Empty arguments would make `go test` run only the repository root package.
    if not selected:
        parser.exit(1, f"No packages selected for {args.suite} shard {args.shard}\n")
    print("\n".join(selected))


if __name__ == "__main__":
    main()
