// Databricks File System (DBFS) API
//
// We recommend using a client created via [databricks.NewWorkspaceClient]
// to simplify the configuration experience.
//
// # Reading and writing files
//
// You can open a file on DBFS for reading or writing with [DbfsAPI.Open].
// This function returns a [Handle] that is compatible with a subset of [io]
// interfaces for reading, writing, and closing.
//
// Uploading a file from an [io.Reader]:
//
//	upload, _ := os.Open("/path/to/local/file.ext")
//	remote, _ := w.Dbfs.Open(ctx, "/path/to/remote/file", files.FileModeWrite|files.FileModeOverwrite)
//	io.Copy(remote, upload)
//	remote.Close()
//
// Downloading a file to an [io.Writer]:
//
//	download, _ := os.Create("/path/to/local")
//	remote, _ := w.Dbfs.Open(ctx, "/path/to/remote/file", files.FileModeRead)
//	_ = io.Copy(download, remote)
//
// # Reading and writing files from buffers
//
// You can read from or write to a DBFS file directly from a byte slice through
// the convenience functions [DbfsAPI.ReadFile] and [DbfsAPI.WriteFile].
//
// Uploading a file from a byte slice:
//
//	buf := []byte("Hello world!")
//	_ = w.Dbfs.WriteFile(ctx, "/path/to/remote/file", buf)
//
// Downloading a file into a byte slice:
//
//	buf, err := w.Dbfs.ReadFile(ctx, "/path/to/remote/file")
//
// # Moving files
//
//	err := w.Dbfs.Move(ctx, files.Move{
//		SourcePath:      "/remote/src/path",
//		DestinationPath: "/remote/dst/path",
//	})
//
// # Creating directories
//
//	w.Dbfs.MkdirsByPath(ctx, "/remote/dir/path")
package files
