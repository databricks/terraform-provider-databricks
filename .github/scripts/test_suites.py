#!/usr/bin/env python3
"""Partition Go tests without downloading dependencies for matrix discovery."""

import argparse
import json
import os
from pathlib import Path
import sys


def replay_suites(root):
    suites = []
    for directory, children, files in os.walk(root / "qa/replay"):
        # Match the directories and filenames ignored by `go list ./...`.
        children[:] = [
            name for name in children
            if name not in {"testdata", "vendor"} and not name.startswith((".", "_"))
        ]
        if any(name.endswith("_test.go") and not name.startswith((".", "_")) for name in files):
            suites.append("./" + Path(directory).relative_to(root).as_posix())
    if not suites:
        raise ValueError("No replay test suites found")
    return sorted(suites)


def replay_matrix(root, shard_count=8):
    suites = replay_suites(root)
    groups = [
        {"shard": index + 1, "packages": []}
        for index in range(min(shard_count, len(suites)))
    ]
    loads = [0] * len(groups)
    # Fixture counts approximate work while keeping new suites automatic.
    weights = {
        suite: max(1, sum(
            len(list((root / suite / "testdata").glob(pattern)))
            for pattern in ["*.yaml", "*.yml"]
        ))
        for suite in suites
    }
    for suite in sorted(suites, key=lambda name: (-weights[name], name)):
        index = min(range(len(groups)), key=lambda index: (loads[index], index))
        groups[index]["packages"].append(suite)
        loads[index] += weights[suite]
    for group in groups:
        group["packages"].sort()
    return groups


def unit_packages(root, packages):
    replay_directories = {(root / suite).resolve() for suite in replay_suites(root)}
    selected = []
    for package in packages:
        import_path, directory = package.rstrip("\n").split(" ", 1)
        if Path(directory).resolve() not in replay_directories:
            selected.append(import_path)
    if not selected:
        raise ValueError("No non-replay Go packages found")
    return selected


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("partition", choices=["matrix", "unit"])
    args = parser.parse_args()
    root = Path(__file__).resolve().parents[2]
    if args.partition == "matrix":
        print(json.dumps({"include": replay_matrix(root)}))
    else:
        print("\n".join(unit_packages(root, sys.stdin)))


if __name__ == "__main__":
    main()
