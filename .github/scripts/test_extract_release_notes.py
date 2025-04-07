#!/usr/bin/env python3

import os
import tempfile
import unittest
from extract_release_notes import extract_release_notes

class TestExtractReleaseNotes(unittest.TestCase):
    def test_extract_release_notes_success(self):
        # Create a temporary changelog file with two releases
        with tempfile.NamedTemporaryFile(mode='w', delete=False) as f:
            f.write("""# Changelog

## Release v1.2.0

- Added new feature X
- Fixed bug Y
- Improved performance Z

## Release v1.1.0

- Initial release
""")
            changelog_path = f.name

        try:
            # Capture stdout
            import io
            import sys
            captured_output = io.StringIO()
            sys.stdout = captured_output

            # Run the function
            result = extract_release_notes(changelog_path)

            # Restore stdout
            sys.stdout = sys.__stdout__

            # Check the results
            self.assertEqual(result, 0)
            expected_output = """- Added new feature X
- Fixed bug Y
- Improved performance Z"""
            self.assertEqual(captured_output.getvalue().strip(), expected_output)

        finally:
            # Clean up the temporary file
            os.unlink(changelog_path)

    def test_extract_release_notes_no_releases(self):
        # Create a temporary changelog file with no releases
        with tempfile.NamedTemporaryFile(mode='w', delete=False) as f:
            f.write("""# Changelog

This is a changelog without any releases.
""")
            changelog_path = f.name

        try:
            # Capture stderr
            import io
            import sys
            captured_output = io.StringIO()
            sys.stderr = captured_output

            # Run the function
            result = extract_release_notes(changelog_path)

            # Restore stderr
            sys.stderr = sys.__stderr__

            # Check the results
            self.assertEqual(result, 1)
            self.assertEqual(captured_output.getvalue().strip(), "No release notes found")

        finally:
            # Clean up the temporary file
            os.unlink(changelog_path)

if __name__ == '__main__':
    unittest.main()