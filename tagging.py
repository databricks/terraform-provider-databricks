#!/usr/bin/env python3

import os
import re
import argparse
from typing import Optional, List, Callable
from dataclasses import dataclass
import subprocess
import time
import json
from datetime import datetime, timezone

NEXT_CHANGELOG_FILE_NAME = "NEXT_CHANGELOG.md"
CHANGELOG_FILE_NAME = "CHANGELOG.md"
PACKAGE_FILE_NAME = ".package.json"


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

    Example:

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

    """

    package: Package
    version: str
    content: str

    def tag_name(self):
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
    with open(filepath, 'r') as file:
        content = json.load(file)
    if "package" in content:
        return content["package"]
    # Legacy SDKs have no packages.
    return ""


def update_version_references(tag_info: TagInfo) -> None:
    """
    Updates the version of the package in code references.
    Code references are defined in .package.json files.
    """

    # Load version patterns from '.package.json' file
    package_file_path = os.path.join(os.getcwd(), tag_info.package.path, PACKAGE_FILE_NAME)
    with open(package_file_path, 'r') as file:
        package_file = json.load(file)

    version = package_file.get('version')
    if not version:
        print(f"Version not found in .package.json. Nothing to update.")
        return

    # Update the versions
    for filename, pattern in version.items():
        loc = os.path.join(os.getcwd(), tag_info.package.path, filename)
        previous_version = re.sub(r'\$VERSION', r"\\d+\\.\\d+\\.\\d+", pattern)
        new_version = re.sub(r'\$VERSION', tag_info.version, pattern)

        with open(loc, 'r') as file:
            content = file.read()

        # Replace the version in the file content
        updated_content = re.sub(previous_version, new_version, content)

        # Write the updated content back to the file
        with open(loc, 'w') as file:
            file.write(updated_content)


def clean_next_changelog(package_path: str) -> None:
    """
    Cleans the "NEXT_CHANGELOG.md" file. It performs 2 operations:
    * Increase the version to the next minor version.
    * Remove release notes. Sections names are kept to
      keep consistency in the section names between releases.
    """

    file_path = os.path.join(os.getcwd(), package_path, NEXT_CHANGELOG_FILE_NAME)
    with open(file_path, 'r') as file:
        content = file.read()

    # Remove content between ### sections
    cleaned_content = re.sub(r'(### [^\n]+\n)\n*([^#]+)', r'\1\n', content)
    # Ensure there is exactly one empty line before each section
    cleaned_content = re.sub(r'(\n*)(###[^\n]+)', r'\n\n\2', cleaned_content)
    # Find the version number
    version_match = re.search(r'Release v(\d+)\.(\d+)\.(\d+)', cleaned_content)
    if not version_match:
        raise Exception("Version not found in the changelog")
    major, minor, patch = map(int, version_match.groups())
    # Prepare next release version.
    # When doing a PR, teams can adjust the version.
    # By default, we increase a minor version, since minor versions releases
    # are more common than patch or major version releases.
    minor += 1
    patch = 0
    new_version = f'Release v{major}.{minor}.{patch}'
    cleaned_content = cleaned_content.replace(version_match.group(0), new_version)

    # Write changes to disk
    with open(file_path, 'w') as file:
        file.write(cleaned_content)


def get_previous_tag_info(package: Package) -> Optional[TagInfo]:
    """
    Extracts the previous tag info from the "CHANGELOG.md" file. 
    Used for failure recovery purposes.
    """
    changelog_path = os.path.join(os.getcwd(), package.path, CHANGELOG_FILE_NAME)

    with open(changelog_path, 'r') as f:
        changelog = f.read()

    # Extract the latest release section using regex
    match = re.search(r"## \[Release\] Release v[\d\.]+.*?(?=\n## (\[Release\] )?Release v|\Z)", changelog, re.S)

    # E.g., for new packages.
    if not match:
        return None

    latest_release = match.group(0)
    version_match = re.search(r'## \[Release\] Release v(\d+\.\d+\.\d+)', latest_release)

    if not version_match:
        raise Exception("Version not found in the changelog")

    return TagInfo(package=package, version=version_match.group(1), content=latest_release)


def get_next_tag_info(package: Package) -> Optional[TagInfo]:
    """
    Extracts the changes from the "NEXT_CHANGELOG.md" file. 
    The result is already processed.
    """
    next_changelog_path = os.path.join(os.getcwd(), package.path, NEXT_CHANGELOG_FILE_NAME)
    # Read NEXT_CHANGELOG.md
    with open(next_changelog_path, 'r') as f:
        next_changelog = f.read()

    # Remove "# NEXT CHANGELOG" line
    next_changelog = re.sub(r'^# NEXT CHANGELOG(\n+)', '', next_changelog, flags=re.MULTILINE)

    # Remove empty sections
    next_changelog = re.sub(r'###[^\n]+\n+(?=##|\Z)', '', next_changelog)
    # Ensure there is exactly one empty line before each section
    next_changelog = re.sub(r'(\n*)(###[^\n]+)', r'\n\n\2', next_changelog)

    if not re.search(r'###', next_changelog):
        print("All sections are empty. No changes will be made to the changelog.")
        return None

    version_match = re.search(r'## Release v(\d+\.\d+\.\d+)', next_changelog)

    if not version_match:
        raise Exception("Version not found in the changelog")

    return TagInfo(package=package, version=version_match.group(1), content=next_changelog)


def write_changelog(tag_info: TagInfo) -> None:
    """
    Updates the changelog with a new tag info.
    """
    changelog_path = os.path.join(os.getcwd(), tag_info.package.path, CHANGELOG_FILE_NAME)
    with open(changelog_path, 'r') as f:
        changelog = f.read()
    updated_changelog = re.sub(r'(# Version changelog\n\n)', f'\\1{tag_info.content.strip()}\n\n\n',
                               changelog)
    with open(changelog_path, 'w') as f:
        f.write(updated_changelog)


def process_package(package: Package) -> TagInfo:
    """
    Processes a package
    """
    # Prepare tag_info from NEXT_CHANGELOG.md
    print(f"Processing package {package.name}")
    tag_info = get_next_tag_info(package)

    # If there are no updates, skip.
    if tag_info is None:
        return

    write_changelog(tag_info)
    clean_next_changelog(package.path)
    update_version_references(tag_info)
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
        result = subprocess.check_output(
            ['git', 'tag', '--list', tag.tag_name()], stderr=subprocess.PIPE, text=True)
        return result.strip() == tag.tag_name()
    except subprocess.CalledProcessError as e:
        # Raise a exception for git command errors
        raise Exception(f"Git command failed: {e.stderr.strip() or e}") from e


def find_last_tags() -> List[TagInfo]:
    """
    Finds the last tags for each package.

    Returns a list of TagInfo objects for each package with a non-None changelog.
    """
    packages = find_packages()

    return [
        info for info in (get_previous_tag_info(package) for package in packages)
        if info is not None
    ]


def find_pending_tags() -> List[TagInfo]:
    """
    Finds all tags that are pending to be applied.
    """
    tag_infos = find_last_tags()
    return [tag for tag in tag_infos if not is_tag_applied(tag)]


def push_changes() -> None:
    """Pushes changes to the remote repository after handling possible merge conflicts."""

    # Create the release metadata file
    file_name = os.path.join(os.getcwd(), ".release_metadata.json")
    metadata = {"timestamp": datetime.now(tz=timezone.UTC).strftime("%Y-%m-%d %H:%M:%S%Z")}
    with open(file_name, "w") as f:
        json.dump(metadata, f, indent=4)

    # Commit the changes
    subprocess.check_output(['git', 'add', '--all'])  # Stage all changes
    subprocess.check_output(['git', 'commit', '-m', 'Release'])  # Commit with message "Release"

    # Push the changes
    subprocess.check_output(['git', 'push'])  # Step 3: Push the commit to the remote


def reset_repository(hash: Optional[str] = None) -> None:
    """
    Reset git to the specified commit. Defaults to HEAD.

    :param hash: The commit hash to reset to. If None, it resets to HEAD.
    """
    # Fetch the latest changes from the remote repository
    subprocess.run(['git', 'fetch'])

    # Determine the commit hash (default to origin/main if none is provided)
    commit_hash = hash or 'origin/main'

    # Construct the Git reset command
    command = ['git', 'reset', '--hard', commit_hash]

    # Execute the git reset command
    subprocess.run(command, check=True)


def retry_function(func: Callable[[], List[TagInfo]],
                   cleanup: Callable[[], None],
                   max_attempts: int = 5,
                   delay: int = 5) -> List[TagInfo]:
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


def update_changelogs(packages: List[Package]) -> List[TagInfo]:
    """
    Updates changelogs and pushes the commits.
    """
    tag_infos = [
        info for info in (process_package(package) for package in packages) if info is not None
    ]
    # If any package was changed, push the changes.
    if tag_infos:
        push_changes()
    return tag_infos


def push_tags(tag_infos: List[TagInfo]) -> None:
    """
    Creates and pushes tags to the repository.
    """
    for tag_info in tag_infos:
        # Create the tag locally
        command_tag = ['git', 'tag', tag_info.tag_name(), '-m', tag_info.content]
        print(f'Running command: {" ".join(command_tag)}')
        subprocess.run(command_tag, check=True)

        # Push the tag to the remote repository
        command_push = ['git', 'push', 'origin', tag_info.tag_name()]
        print(f'Running command: {" ".join(command_push)}')
        subprocess.run(command_push, check=True)


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
        ['git', 'log', '-n', '1', '--format=%H', '--', '.release_metadata.json'],
        text=True).strip()

    # If no commit is found, raise an exception
    if not commit_hash:
        raise ValueError("No commit found for .release_metadata.json")

    # Reset the repository to the commit
    reset_repository(commit_hash)


def get_package_from_args() -> Optional[str]:
    """
    Retrieves an optional package
    python3 ./tagging.py --package <name>
    """
    parser = argparse.ArgumentParser(description='Update changelogs and tag the release.')
    parser.add_argument('--package', '-p', type=str, help='Tag a single package')
    args = parser.parse_args()
    return args.package


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
    package_name = get_package_from_args()
    pending_tags = find_pending_tags()

    # pending_tags is non-empty only when the tagging process previously failed or interrupted.
    # We must complete the interrupted tagging process before starting a new one to avoid inconsistent states and missing changelog entries.
    # Therefore, we don't support specifying the package until the previously started process has been successfully completed.
    if pending_tags and package_name:
        pending_packages = [tag.package.name for tag in pending_tags]
        raise Exception(
            f"Cannot release package {package_name}. Pending release for {pending_packages}")

    if pending_tags:
        print("Found pending tags from previous executions, entering recovery mode.")
        pull_last_release_commit()
        push_tags(pending_tags)
        return

    packages = find_packages()
    # If a package is specified as an argument, only process that package
    if package_name:
        packages = [package for package in packages if package.name == package_name]

    pending_tags = retry_function(
        func=lambda: update_changelogs(packages), cleanup=reset_repository)
    push_tags(pending_tags)


if __name__ == "__main__":
    process()
