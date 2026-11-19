# Local replay inputs

These are deliberately small, shared input files. Python files contain a single
comment; the ZIP contains one comment-only script. The two notebook formats keep
one cell/command; dashboard JSON contains one page. Updated variants differ in
content so lifecycle tests exercise content changes without large artifacts.

DBC archives contain `manifest.mf` and a `NotebookV1` notebook; timestamps are fixed.
Empty-file scenarios use the zero-byte `empty` fixture.
