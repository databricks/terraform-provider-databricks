#!/usr/bin/env python3
import os
import re
import sys
import glob
import subprocess


def load_ignores():
    ignores = set()
    fail = False
    for ind, line in enumerate(open(".wsignore")):
        line = line.strip()
        if not line:
            continue
        if line.startswith("#"):
            continue
        expanded = glob.glob(line, recursive=True)
        if len(expanded) == 0:
            print(f".wsignore:{ind + 1}: No matches for line: {line}")
            fail = True
        ignores.update(expanded)
    if fail:
        sys.exit(1)
    return ignores


def count_trailing_newlines(s):
    match = re.search(r"(\n+)$", s)
    return len(match.group(1)) if match else 0


def validate_contents(data):
    if not data:
        return
    try:
        text = data.decode("utf")
    except Exception as ex:
        yield f" Failed to decode utf-8: {ex}"
        return

    for i, line in enumerate(text.split("\n")):
        if not line:
            continue
        if line.strip() == "":
            yield f"{i + 1}: Whitespace-only line"
            continue
        if line.rstrip() != line:
            yield f"{i + 1}: Trailing whitespace {line[-200:]!r}"

    newlines = count_trailing_newlines(text)

    if newlines == 0:
        yield " File does not end with a newline"

    if newlines >= 2:
        yield f" {newlines} newlines at the end"


def main():
    quiet = "-q" in sys.argv
    files = subprocess.check_output(["git", "ls-files"], encoding="utf-8").split()
    ignores = load_ignores()
    n_checked = 0
    n_skipped = 0
    n_errored = 0
    for f in files:
        if not os.path.isfile(f):
            n_skipped += 1
            continue
        if f in ignores:
            n_skipped += 1
            continue
        with open(f, "rb") as file:
            data = file.read()
        error = False
        for msg in validate_contents(data):
            print(f"{f}:{msg}")
            error = True
        n_checked += 1
        n_errored += 1 if error else 0

    if not quiet:
        sys.stderr.write(f"{n_checked} checked, {n_skipped} skipped, {n_errored} failed.\n")
    sys.exit(1 if n_errored else 0)


if __name__ == "__main__":
    try:
        main()
    except BrokenPipeError:
        sys.exit(1)
