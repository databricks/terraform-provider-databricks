#!/usr/bin/env python3

import re
import sys

def extract_release_notes(changelog_path):
    """Extract the content between the first two level-2 headings from CHANGELOG.md."""
    try:
        with open(changelog_path, 'r') as f:
            content = f.read()
            
        # Find the first two level-2 headings
        pattern = r'^## Release v[\d.]+\n\n(.*?)(?=^## Release v[\d.]+|\Z)'
        match = re.search(pattern, content, re.MULTILINE | re.DOTALL)
        
        if match:
            release_notes = match.group(1).strip()
            print(release_notes)
            return 0
        else:
            print("No release notes found", file=sys.stderr)
            return 1
            
    except Exception as e:
        print(f"Error: {str(e)}", file=sys.stderr)
        return 1

if __name__ == "__main__":
    changelog_path = "CHANGELOG.md"
    sys.exit(extract_release_notes(changelog_path)) 