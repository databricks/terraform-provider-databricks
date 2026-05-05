#!/usr/bin/env python3
# /// script
# dependencies = ["PyGithub>=2,<3", "pyjwt<2.12.0", "charset-normalizer<3.4.6"]
# ///

import os
import re
import argparse
from typing import Optional, List, Callable, Dict
from dataclasses import dataclass, replace
import subprocess
import time
import json
from github import Github, Repository, InputGitTreeElement, InputGitAuthor
from datetime import datetime, timezone

NEXT_CHANGELOG_FILE_NAME = "NEXT_CHANGELOG.md"
CHANGELOG_FILE_NAME = "CHANGELOG.md"
PACKAGE_FILE_NAME = ".package.json"
CODEGEN_FILE_NAME = ".codegen.json"
CREATED_TAGS_FILE_NAME = "created_tags.json"
"""
This script tags the release of the SDKs using a combination of the GitHub API and Git commands.
It reads the local repository to determine necessary changes, updates changelogs, and creates tags.

### How it Works:
- It does **not** modify the local repository directly.
- Instead of committing and pushing changes locally, it uses the **GitHub API** to create commits and tags.
"""


@dataclass(frozen=True)
class Version:
    """
    A semver 2.0.0-compliant version (https://semver.org).

    Mirrors the API of the `semver` PyPI package so this implementation can be
    swapped for that library if it is ever added to the wheelhouse. Supports
    parsing, stringification, and the two bumps we need: minor (for stable
    releases) and prerelease (for release trains).
    """

    # Permissive pattern for locating a semver version string inside larger
    # text (e.g. a changelog header). Callers use it in f-strings; strict
    # validation happens via Version.parse.
    PATTERN = r"\d+\.\d+\.\d+(?:-[0-9A-Za-z.-]+)?(?:\+[0-9A-Za-z.-]+)?"

    # Strict anchored regex per https://semver.org. Rejects leading zeros in
    # numeric identifiers and invalid pre-release/build identifier charsets.
    _PARSE_REGEX = re.compile(
        r"^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)"
        r"(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)"
        r"(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?"
        r"(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$"
    )

    major: int
    minor: int
    patch: int
    prerelease: str = ""
    build: str = ""

    @classmethod
    def parse(cls, text: str) -> "Version":
        """Parse a semver string, raising ValueError on malformed input."""
        match = cls._PARSE_REGEX.match(text)
        if not match:
            raise ValueError(f"Invalid semver version: {text!r}")
        major, minor, patch, prerelease, build = match.groups()
        return cls(
            major=int(major),
            minor=int(minor),
            patch=int(patch),
            prerelease=prerelease or "",
            build=build or "",
        )

    def __str__(self) -> str:
        result = f"{self.major}.{self.minor}.{self.patch}"
        if self.prerelease:
            result += f"-{self.prerelease}"
        if self.build:
            result += f"+{self.build}"
        return result

    def bump_minor(self) -> "Version":
        """
        Bump the minor version and reset patch.

        Per semver item 9, a pre-release version has lower precedence than
        the same MAJOR.MINOR.PATCH, so bumping to a new minor drops any
        pre-release and build metadata.
        """
        return Version(major=self.major, minor=self.minor + 1, patch=0)

    def bump_prerelease(self) -> "Version":
        """
        Increment the rightmost numeric identifier in the pre-release.

        Matches the npm `prerelease` bump semantics:
            0.0.0-alpha.1 -> 0.0.0-alpha.2
            0.0.0-alpha   -> 0.0.0-alpha.1
            0.0.0-rc.1.2  -> 0.0.0-rc.1.3

        Raises ValueError if the version has no pre-release to bump.
        Build metadata is dropped since it does not affect precedence.
        """
        if not self.prerelease:
            raise ValueError(f"Cannot bump prerelease of {self}: no pre-release component")
        parts = self.prerelease.split(".")
        for i in range(len(parts) - 1, -1, -1):
            if parts[i].isdigit():
                parts[i] = str(int(parts[i]) + 1)
                return replace(self, prerelease=".".join(parts), build="")
        # No numeric identifier exists; append ".1" to start a counter.
        return replace(self, prerelease=f"{self.prerelease}.1", build="")

    def next_release_version(self) -> "Version":
        """
        Default next version for the changelog after this one is released.

        If on a pre-release track, stay on it by bumping the pre-release
        identifier (npm convention). Otherwise, bump the minor version,
        the script's historical default for stable releases. Teams can
        override the default in the release PR.
        """
        if self.prerelease:
            return self.bump_prerelease()
        return self.bump_minor()


# GitHub does not support signing commits for GitHub Apps directly.
# This class replaces usages for git commands such as "git add", "git commit", and "git push".
@dataclass
class GitHubRepo:
    def __init__(self, repo: Repository):
        self.repo = repo
        self.changed_files: list[InputGitTreeElement] = []
        self.ref = "heads/main"
        head_ref = self.repo.get_git_ref(self.ref)
        self.sha = head_ref.object.sha

    # Replaces "git add file"
    def add_file(self, loc: str, content: str):
        local_path = os.path.relpath(loc, os.getcwd())
        print(f"Adding file {local_path}")
        blob = self.repo.create_git_blob(content=content, encoding="utf-8")
        element = InputGitTreeElement(path=local_path, mode="100644", type="blob", sha=blob.sha)
        self.changed_files.append(element)

    # Replaces "git commit && git push"
    def commit_and_push(self, message: str):
        head_ref = self.repo.get_git_ref(self.ref)
        base_tree = self.repo.get_git_tree(sha=head_ref.object.sha)
        new_tree = self.repo.create_git_tree(self.changed_files, base_tree)
        parent_commit = self.repo.get_git_commit(head_ref.object.sha)

        new_commit = self.repo.create_git_commit(message=message, tree=new_tree, parents=[parent_commit])
        # Update branch reference
        head_ref.edit(new_commit.sha)
        self.sha = new_commit.sha

    def reset(self, sha: Optional[str] = None):
        self.changed_files = []
        if sha:
            self.sha = sha
        else:
            head_ref = self.repo.get_git_ref(self.ref)
            self.sha = head_ref.object.sha

    def tag(self, tag_name: str, tag_message: str):
        # Create a tag pointing to the new commit
        # The email MUST be the GitHub Apps email.
        # Otherwise, the tag will not be verified.
        tagger = InputGitAuthor(
            name="Databricks SDK Release Bot", email="DECO-SDK-Tagging[bot]@users.noreply.github.com"
        )

        tag = self.repo.create_git_tag(tag=tag_name, message=tag_message, object=self.sha, type="commit", tagger=tagger)
        # Create a Git ref (the actual reference for the tag in the repo)
        self.repo.create_git_ref(ref=f"refs/tags/{tag_name}", sha=tag.sha)


gh: Optional[GitHubRepo] = None


@dataclass
class Package:
    """
    Represents a package in the repository.
    :name: The package name.
    :path: The path to the package relative to the repository root.
    """

    name: str
    path: str


@dataclass
class TagInfo:
    """
    Represents all changes on a release.
    :package: package info.
    :version: release version for the package. Format: v<major>.<minor>.<pacth>
    :content: changes for the release, as they appear in the changelog.
              When written to CHANGELOG.md, the current date (YYYY-MM-DD) is automatically added.

    Example (from NEXT_CHANGELOG.md):

    ## Release v0.56.0

    ### New Features and Improvements
    * Feature
    * Some improvement

    ### Bug Fixes
    * Bug fix

    ### Documentation
    * Doc Changes

    ### Internal Changes
    * More Changes

    ### API Changes
    * Add new Service

    Note: When written to CHANGELOG.md, the header becomes: ## Release v0.56.0 (YYYY-MM-DD)

    """

    package: Package
    version: str
    content: str

    def tag_name(self) -> str:
        return f"{self.package.name}/v{self.version}" if self.package.name else f"v{self.version}"


def get_package_name(package_path: str) -> str:
    """
    Returns the package name from the package path.
    The name is found inside the .package.json file:
    {
        "package": "package_name"
    }
    """
    filepath = os.path.join(os.getcwd(), package_path, PACKAGE_FILE_NAME)
    with open(filepath, "r") as file:
        content = json.load(file)
    if "package" in content:
        return content["package"]
    # Legacy SDKs have no packages.
    return ""


def stage_version_updates(tag_infos: List[TagInfo], packages: List[Package]) -> None:
    """
    Stages all version-related edits for the release in a single pass over
    every package the workspace already opts in via ``.package.json``.
    """

    # Load patterns from '.codegen.json' at the top level of the repository.
    package_file_path = os.path.join(os.getcwd(), CODEGEN_FILE_NAME)
    with open(package_file_path, "r") as file:
        codegen = json.load(file)

    version_patterns = codegen.get("version", {})
    dep_patterns = codegen.get("dependency_pattern", {})
    name_template = codegen.get("dependency_name_template", "")

    if not version_patterns and not dep_patterns:
        print("Neither `version` nor `dependency_pattern` found in .codegen.json. Nothing to update.")
        return

    bumped_by_dir: Dict[str, TagInfo] = {info.package.path: info for info in tag_infos}
    new_dep_versions = compute_dependency_rewrites(tag_infos, name_template)

    files = sorted(set(version_patterns.keys()) | set(dep_patterns.keys()))

    for pkg in packages:
        for filename in files:
            loc = os.path.join(os.getcwd(), pkg.path, filename)

            with open(loc, "r") as file:
                content = file.read()
            original = content

            # Own version (only when this package is being released and the
            # file has a version pattern declared).
            info = bumped_by_dir.get(pkg.path)
            if info is not None and filename in version_patterns:
                pattern = version_patterns[filename]
                previous_version = pattern.replace("$VERSION", Version.PATTERN)
                new_version = pattern.replace("$VERSION", info.version)
                content = re.sub(previous_version, new_version, content)

            # Sibling dependency rewrites (only when the file has a
            # dependency pattern and there is at least one bumped sibling).
            if filename in dep_patterns and new_dep_versions:
                content = rewrite_dependencies(content, dep_patterns[filename], new_dep_versions)

            if content != original:
                gh.add_file(loc, content)


def compute_dependency_rewrites(
    tag_infos: List[TagInfo],
    name_template: str,
) -> Dict[str, str]:
    """
    Returns a map of dependency-name to the new semver string for each
    bumped package.
    """
    if not name_template:
        return {}
    rewrites: Dict[str, str] = {}
    for info in tag_infos:
        # Skip legacy releases that don't have a per-package name; their
        # tag_info has an empty package.name and they can't be referenced
        # as a sibling dep anyway.
        if not info.package.name:
            continue
        dep_name = name_template.replace("$PACKAGE", info.package.name)
        rewrites[dep_name] = info.version
    return rewrites


def rewrite_dependencies(content: str, pattern: str, new_versions: Dict[str, str]) -> str:
    """
    Apply ``pattern`` (with ``$DEPENDENCY`` and ``$VERSION`` placeholders) to
    rewrite every entry in ``content`` whose dependency name appears in
    ``new_versions``.
    """
    # Sentinel strings used to protect the placeholders through re.escape:
    # we substitute them in, escape the whole template, then swap them out
    # for the dep-name literal and Version.PATTERN. Control characters so
    # they can't collide with anything in real .codegen.json patterns.
    dep_sentinel = "\x01DEPENDENCY\x01"
    ver_sentinel = "\x01VERSION\x01"

    for dep_name, new_value in new_versions.items():
        regex = pattern.replace("$DEPENDENCY", dep_sentinel).replace("$VERSION", ver_sentinel)
        regex = re.escape(regex)
        regex = regex.replace(re.escape(dep_sentinel), re.escape(dep_name))
        regex = regex.replace(re.escape(ver_sentinel), Version.PATTERN)

        # Build the literal replacement text by substituting the same
        # placeholders directly. A lambda is used instead of a string to
        # avoid re.sub interpreting \1, \g<...>, etc. inside the value.
        replacement_text = pattern.replace("$DEPENDENCY", dep_name).replace("$VERSION", new_value)
        content = re.sub(regex, lambda _m, text=replacement_text: text, content)
    return content


def clean_next_changelog(package_path: str) -> None:
    """
    Cleans the "NEXT_CHANGELOG.md" file. It performs 2 operations:
    * Increase the version to the next minor version.
    * Remove release notes. Sections names are kept to
      keep consistency in the section names between releases.
    """

    file_path = os.path.join(os.getcwd(), package_path, NEXT_CHANGELOG_FILE_NAME)
    with open(file_path, "r") as file:
        content = file.read()

    # Remove content between ### sections.
    cleaned_content = re.sub(r"(### [^\n]+\n)(?:.*?\n?)*?(?=###|$)", r"\1", content)
    # Ensure there is exactly one empty line before each section.
    cleaned_content = re.sub(r"(\n*)(###[^\n]+)", r"\n\n\2", cleaned_content)
    # Find the version number and compute the default next release version.
    # Teams can adjust the version in the PR if the default is not desired.
    # For stable versions, bump minor (historical default since minor releases
    # are more common than patch or major). For pre-release versions, stay on
    # the same track by bumping the pre-release identifier (npm convention).
    version_match = re.search(rf"Release v({Version.PATTERN})", cleaned_content)
    if not version_match:
        raise Exception("Version not found in the changelog")
    current = Version.parse(version_match.group(1))
    new_header = f"Release v{current.next_release_version()}"
    cleaned_content = cleaned_content.replace(version_match.group(0), new_header)

    # Update file with cleaned content
    gh.add_file(file_path, cleaned_content)


def get_previous_tag_info(package: Package) -> Optional[TagInfo]:
    """
    Extracts the previous tag info from the "CHANGELOG.md" file.
    Used for failure recovery purposes.
    """
    changelog_path = os.path.join(os.getcwd(), package.path, CHANGELOG_FILE_NAME)

    with open(changelog_path, "r") as f:
        changelog = f.read()

    # Extract the latest release section using regex.
    match = re.search(
        rf"## (\[Release\] )?Release v{Version.PATTERN}.*?(?=\n## (\[Release\] )?Release v|\Z)",
        changelog,
        re.S,
    )

    # E.g., for new packages.
    if not match:
        return None

    latest_release = match.group(0)
    version_match = re.search(rf"## (\[Release\] )?Release v({Version.PATTERN})", latest_release)

    if not version_match:
        raise Exception("Version not found in the changelog")

    # Validate the extracted string is spec-compliant; fail loudly otherwise.
    version = str(Version.parse(version_match.group(2)))
    return TagInfo(package=package, version=version, content=latest_release)


def get_next_tag_info(package: Package) -> Optional[TagInfo]:
    """
    Extracts the changes from the "NEXT_CHANGELOG.md" file.
    The result is already processed.
    """
    next_changelog_path = os.path.join(os.getcwd(), package.path, NEXT_CHANGELOG_FILE_NAME)
    # Read NEXT_CHANGELOG.md
    with open(next_changelog_path, "r") as f:
        next_changelog = f.read()

    # Remove "# NEXT CHANGELOG" line
    next_changelog = re.sub(r"^# NEXT CHANGELOG(\n+)", "", next_changelog, flags=re.MULTILINE)

    # Remove empty sections
    next_changelog = re.sub(r"###[^\n]+\n+(?=##|\Z)", "", next_changelog)
    # Ensure there is exactly one empty line before each section
    next_changelog = re.sub(r"(\n*)(###[^\n]+)", r"\n\n\2", next_changelog)

    if not re.search(r"###", next_changelog):
        print("All sections are empty. No changes will be made to the changelog.")
        return None

    version_match = re.search(rf"## Release v({Version.PATTERN})", next_changelog)

    if not version_match:
        raise Exception("Version not found in the changelog")

    # Validate the extracted string is spec-compliant; fail loudly otherwise.
    version = str(Version.parse(version_match.group(1)))
    return TagInfo(package=package, version=version, content=next_changelog)


def write_changelog(tag_info: TagInfo) -> None:
    """
    Updates the changelog with a new tag info.
    """
    changelog_path = os.path.join(os.getcwd(), tag_info.package.path, CHANGELOG_FILE_NAME)
    with open(changelog_path, "r") as f:
        changelog = f.read()

    # Add current date to the release header.
    current_date = datetime.now(tz=timezone.utc).strftime("%Y-%m-%d")
    content_with_date = re.sub(
        rf"## Release v({Version.PATTERN})",
        rf"## Release v\1 ({current_date})",
        tag_info.content.strip(),
    )

    updated_changelog = re.sub(r"(# Version changelog\n\n)", f"\\1{content_with_date}\n\n\n", changelog)
    gh.add_file(changelog_path, updated_changelog)


def process_package(package: Package) -> TagInfo:
    """
    Processes a package's changelog scaffolding for the release.
    """
    print(f"Processing package {package.name}")
    tag_info = get_next_tag_info(package)

    # If there are no updates, skip.
    if tag_info is None:
        return

    write_changelog(tag_info)
    clean_next_changelog(package.path)
    return tag_info


def find_packages() -> List[Package]:
    """
    Returns all directories which contains a ".package.json" file.
    """
    paths = _find_directories_with_file(PACKAGE_FILE_NAME)
    return [Package(name=get_package_name(path), path=path) for path in paths]


def _find_directories_with_file(target_file: str) -> List[str]:
    root_path = os.getcwd()
    matching_directories = []

    for dirpath, _, filenames in os.walk(root_path):
        if target_file in filenames:
            path = os.path.relpath(dirpath, root_path)
            # If the path is the root directory (e.g., SDK V0), set it to an empty string.
            if path == ".":
                path = ""
            matching_directories.append(path)

    return matching_directories


def is_tag_applied(tag: TagInfo) -> bool:
    """
    Returns whether a tag is already applied in the repository.

    :param tag: The tag to check.
    :return: True if the tag is applied, False otherwise.
    :raises Exception: If the git command fails.
    """
    try:
        # Check if the specific tag exists
        result = subprocess.check_output(["git", "tag", "--list", tag.tag_name()], stderr=subprocess.PIPE, text=True)
        return result.strip() == tag.tag_name()
    except subprocess.CalledProcessError as e:
        # Raise a exception for git command errors
        raise Exception(f"Git command failed: {e.stderr.strip() or e}") from e


def find_last_release_tag(package: Package) -> Optional[str]:
    """
    Returns the most recent ``<package>/v*`` tag in the repository, or
    ``None`` if no such tag exists. Tags are sorted by semver ordering
    (``--sort=-v:refname``) so pre-releases sort below their stable
    counterparts.

    :raises Exception: If the git command fails.
    """
    pattern = f"{package.name}/v*" if package.name else "v*"
    try:
        output = subprocess.check_output(
            ["git", "tag", "--list", pattern, "--sort=-v:refname"],
            stderr=subprocess.PIPE,
            text=True,
        ).strip()
    except subprocess.CalledProcessError as e:
        raise Exception(f"Git command failed: {e.stderr.strip() or e}") from e
    if not output:
        return None
    return output.split("\n")[0].strip()


def has_commits_since_tag(tag: str, path: str) -> bool:
    """
    Returns True iff at least one commit reachable from HEAD but not from
    ``tag`` touches ``path``. Used to detect that a sibling dependency has
    unreleased changes that would ship stale if we tagged a dependent
    without re-tagging the dependency.

    :raises Exception: If the git command fails.
    """
    args = ["git", "log", "--oneline", f"{tag}..HEAD", "--", path or "."]
    try:
        output = subprocess.check_output(args, stderr=subprocess.PIPE, text=True).strip()
    except subprocess.CalledProcessError as e:
        raise Exception(f"Git command failed: {e.stderr.strip() or e}") from e
    return bool(output)


def check_dependency_freshness(tag_infos: List[TagInfo], all_packages: List[Package]) -> None:
    """
    Hard-fails when a package being released depends on a sibling package
    that has unreleased commits since its last tag.

    Why: dependency rewrites (``stage_version_updates``) only fire for
    siblings that are *also* being released. Without this check, releasing
    package_a alone — when package_b has commits since its last tag —
    publishes ``package_a@new`` pinning the *old* package_b artifact, which
    won't have the changes package_a's source depends on. The check is
    commit-based (not changelog-based) so a missing ``NEXT_CHANGELOG.md``
    entry on package_b is still caught.

    No-op when ``.codegen.json`` declares no dependency pattern (legacy
    SDKs without per-package wiring).
    """
    if not tag_infos:
        return

    package_file_path = os.path.join(os.getcwd(), CODEGEN_FILE_NAME)
    with open(package_file_path, "r") as file:
        codegen = json.load(file)

    name_template = codegen.get("dependency_name_template", "")
    dep_patterns = codegen.get("dependency_pattern", {})
    if not name_template or not dep_patterns:
        return

    releasing_paths = {info.package.path for info in tag_infos}
    by_dep_name: Dict[str, Package] = {}
    for pkg in all_packages:
        if not pkg.name:
            continue
        by_dep_name[name_template.replace("$PACKAGE", pkg.name)] = pkg

    issues: List[str] = []
    for info in tag_infos:
        for filename, pattern in dep_patterns.items():
            loc = os.path.join(os.getcwd(), info.package.path, filename)
            if not os.path.exists(loc):
                continue
            with open(loc, "r") as f:
                content = f.read()

            for dep_name, dep_pkg in by_dep_name.items():
                if dep_pkg.path == info.package.path:
                    continue
                if dep_pkg.path in releasing_paths:
                    continue

                # Same regex construction used by ``rewrite_dependencies``,
                # so "is this dep referenced?" matches "would the rewrite
                # touch it?". Keeps the two in lockstep.
                regex = (
                    re.escape(pattern)
                    .replace(re.escape("$DEPENDENCY"), re.escape(dep_name))
                    .replace(re.escape("$VERSION"), Version.PATTERN)
                )
                if not re.search(regex, content):
                    continue

                last_tag = find_last_release_tag(dep_pkg)
                if last_tag is None:
                    # No prior tag means the dep was never released; we
                    # can't reason about staleness. Surface it anyway so
                    # the human resolves it explicitly.
                    issues.append(
                        f"{info.package.name} depends on {dep_pkg.name}, "
                        f"which has never been released. Release "
                        f"{dep_pkg.name} first or include it in this run."
                    )
                    continue
                if has_commits_since_tag(last_tag, dep_pkg.path):
                    issues.append(
                        f"{info.package.name} depends on {dep_pkg.name}, "
                        f"which has commits since {last_tag} but is not "
                        f"being released. Either release {dep_pkg.name} "
                        f"as well, or hold this release until its changes "
                        f"are reverted."
                    )

    if issues:
        raise Exception("Dependency freshness check failed:\n  - " + "\n  - ".join(issues))


def find_last_tags() -> List[TagInfo]:
    """
    Finds the last tags for each package.

    Returns a list of TagInfo objects for each package with a non-None changelog.
    """
    packages = find_packages()

    return [info for info in (get_previous_tag_info(package) for package in packages) if info is not None]


def find_pending_tags() -> List[TagInfo]:
    """
    Finds all tags that are pending to be applied.
    """
    tag_infos = find_last_tags()
    return [tag for tag in tag_infos if not is_tag_applied(tag)]


def generate_commit_message(tag_infos: List[TagInfo]) -> str:
    """
    Generates a commit message for the release.
    """
    if not tag_infos:
        raise Exception("No tag infos provided to generate commit message")

    info = tag_infos[0]
    # Legacy mode for SDKs without per service packaging
    if not info.package.name:
        if len(tag_infos) > 1:
            raise Exception("Multiple packages found in legacy mode")
        return f"[Release] Release v{info.version}\n\n{info.content}"

    # Sort tag_infos by package name for consistency
    tag_infos.sort(key=lambda info: info.package.name)
    return "Release\n\n" + "\n\n".join(
        f"## {info.package.name}/v{info.version}\n\n{info.content}" for info in tag_infos
    )


def push_changes(tag_infos: List[TagInfo]) -> None:
    """Pushes changes to the remote repository after handling possible merge conflicts."""

    commit_message = generate_commit_message(tag_infos)

    # Create the release metadata file
    file_name = os.path.join(os.getcwd(), ".release_metadata.json")
    metadata = {"timestamp": datetime.now(tz=timezone.utc).strftime("%Y-%m-%d %H:%M:%S%z")}
    content = json.dumps(metadata, indent=4)
    gh.add_file(file_name, content)

    gh.commit_and_push(commit_message)


def reset_repository(hash: Optional[str] = None) -> None:
    """
    Reset git to the specified commit. Defaults to HEAD.

    :param hash: The commit hash to reset to. If None, it resets to HEAD.
    """
    # Fetch the latest changes from the remote repository
    subprocess.run(["git", "fetch"])

    # Determine the commit hash (default to origin/main if none is provided)
    commit_hash = hash or "origin/main"

    # Reset in memory changed files and the commit hash
    gh.reset(hash)

    # Construct the Git reset command
    command = ["git", "reset", "--hard", commit_hash]

    # Execute the git reset command
    subprocess.run(command, check=True)


def retry_function(
    func: Callable[[], List[TagInfo]], cleanup: Callable[[], None], max_attempts: int = 5, delay: int = 5
) -> List[TagInfo]:
    """
    Calls a function call up to `max_attempts` times if an exception occurs.

    :param func: The function to call.
    :param cleanup: Cleanup function in between retries
    :param max_attempts: The maximum number of retries.
    :param delay: The delay between retries in seconds.
    :return: The return value of the function, or None if all retries fail.
    """
    attempts = 0
    while attempts <= max_attempts:
        try:
            return func()  # Call the function
        except Exception as e:
            attempts += 1
            print(f"Attempt {attempts} failed: {e}")
            if attempts < max_attempts:
                time.sleep(delay)  # Wait before retrying
                cleanup()
            else:
                print("All retry attempts failed.")
                raise e  # Re-raise the exception after max retries


def update_changelogs(selected_packages: List[Package], all_packages: List[Package]) -> List[TagInfo]:
    """
    Updates changelogs and pushes the commits.

    ``selected_packages`` are the packages whose ``NEXT_CHANGELOG.md`` is
    consulted to decide what gets released this run. ``all_packages`` is
    the full repo inventory used for cross-package dep rewrites.

    The freshness check is deliberately *not* called here. ``process``
    runs it before entering the retry loop so a freshness violation
    fails fast — the check is deterministic against the same git state,
    so wrapping it in retry would just delay the same failure five
    times.
    """
    tag_infos = [info for info in (process_package(package) for package in selected_packages) if info is not None]
    # If any package was changed, stage version updates and push.
    if tag_infos:
        stage_version_updates(tag_infos, all_packages)
        push_changes(tag_infos)
    return tag_infos


def preview_tag_infos(packages: List[Package]) -> List[TagInfo]:
    """
    Read-only sibling of ``process_package``: returns the TagInfos that
    would be released for ``packages`` without writing any changelog
    edits. ``process`` calls this before the retry loop so the freshness
    check has a snapshot to validate against. ``process_package`` will
    re-derive the same TagInfos when ``update_changelogs`` runs; the
    duplication is just a couple of NEXT_CHANGELOG.md reads.
    """
    return [info for info in (get_next_tag_info(package) for package in packages) if info is not None]


def push_tags(tag_infos: List[TagInfo]) -> None:
    """
    Creates and pushes tags to the repository.

    As a side effect, writes the names of successfully created tags to
    ``./created_tags.json`` so that workflows triggering this script can
    discover what was produced (the GitHub Actions workflow uploads this
    file as the ``created-tags`` artifact).

    Schema:
        {"tags": ["service-a/v1.2.3", "service-b/v0.4.0"]}

    The manifest is written even if tag creation fails partway through:
    tags that succeeded before the failure are flushed before the
    exception is re-raised, so recovery-mode runs still surface their
    output.
    """
    created: List[str] = []
    try:
        for tag_info in tag_infos:
            gh.tag(tag_info.tag_name(), tag_info.content)
            created.append(tag_info.tag_name())
    finally:
        manifest_path = os.path.join(os.getcwd(), CREATED_TAGS_FILE_NAME)
        with open(manifest_path, "w") as f:
            json.dump({"tags": created}, f)


def run_command(command: List[str]) -> str:
    """
    Runs a command and returns the output
    """
    output = subprocess.check_output(command)
    print(f'Running command: {" ".join(command)}')
    return output.decode()


def pull_last_release_commit() -> None:
    """
    Reset the repository to the last release.
    Uses commit for last change to .release_metadata.json, since it's only updated on releases.
    """
    commit_hash = subprocess.check_output(
        ["git", "log", "-n", "1", "--format=%H", "--", ".release_metadata.json"], text=True
    ).strip()

    # If no commit is found, raise an exception
    if not commit_hash:
        raise ValueError("No commit found for .release_metadata.json")

    # Reset the repository to the commit
    reset_repository(commit_hash)


def get_packages_from_args() -> List[str]:
    """
    Retrieves the list of packages to tag.

        python3 ./tagging.py --package <name>              # single package
        python3 ./tagging.py --package <name1>,<name2>     # multiple packages

    Returns an empty list when --package is omitted, which means all packages
    with pending releases will be tagged.
    """
    parser = argparse.ArgumentParser(description="Update changelogs and tag the release.")
    parser.add_argument(
        "--package",
        "-p",
        type=str,
        default="",
        help="Comma-separated list of packages to tag. Leave empty to tag all packages with pending releases.",
    )
    args = parser.parse_args()
    return [name.strip() for name in args.package.split(",") if name.strip()]


def init_github():
    token = os.environ["GITHUB_TOKEN"]
    repo_name = os.environ["GITHUB_REPOSITORY"]
    g = Github(token)
    repo = g.get_repo(repo_name)
    global gh
    gh = GitHubRepo(repo)


def process():
    """
    Main entry point for tagging process.

    Tagging process consist of multiple steps:
    * For each package, update the corresponding CHANGELOG.md file based on the contents of NEXT_CHANGELOG.md file
    * If any package has been updated, commit and push the changes.
    * Apply and push the new tags matching the version.

    If a specific pagkage is provided as a parameter, only that package will be tagged.

    If any tag are pending from an early process, it will skip updating the CHANGELOG.md files and only apply the tags.
    """

    package_names = get_packages_from_args()
    pending_tags = find_pending_tags()

    # pending_tags is non-empty only when the tagging process previously failed or interrupted.
    # We must complete the interrupted tagging process before starting a new one to avoid inconsistent states and missing changelog entries.
    # Therefore, we don't support specifying packages until the previously started process has been successfully completed.
    if pending_tags and package_names:
        pending_packages = [tag.package.name for tag in pending_tags]
        raise Exception(f"Cannot release packages {package_names}. Pending release for {pending_packages}")

    if pending_tags:
        print("Found pending tags from previous executions, entering recovery mode.")
        pull_last_release_commit()
        push_tags(pending_tags)
        return

    all_packages = find_packages()
    # If packages are specified as an argument, only release those — but
    # dep rewrites and the freshness check still operate over the full
    # set.
    selected_packages = all_packages
    if package_names:
        selected_packages = [package for package in all_packages if package.name in package_names]

    # Run the freshness check against a read-only preview before the
    # retry loop, since the check is deterministic. A freshness
    # violation fails the run immediately, with no commits, no tags, no
    # retry storm.
    check_dependency_freshness(preview_tag_infos(selected_packages), all_packages)

    pending_tags = retry_function(
        func=lambda: update_changelogs(selected_packages, all_packages),
        cleanup=reset_repository,
    )
    push_tags(pending_tags)


def validate_git_root():
    """
    Validate that the script is run from the root of the repository.
    """
    repo_root = subprocess.check_output(["git", "rev-parse", "--show-toplevel"]).strip().decode("utf-8")
    current_dir = subprocess.check_output(["pwd"]).strip().decode("utf-8")
    if repo_root != current_dir:
        raise Exception("Please run this script from the root of the repository.")


if __name__ == "__main__":
    validate_git_root()
    init_github()
    process()
