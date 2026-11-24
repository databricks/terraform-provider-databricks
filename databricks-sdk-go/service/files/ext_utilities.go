package files

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// FileMode conveys user intent when opening a file.
type FileMode int

const (
	// Exactly one of FileModeRead or FileModeWrite must be specified.
	FileModeRead FileMode = 1 << iota
	FileModeWrite
	FileModeOverwrite
)

// Maximum read or write length for the DBFS API.
const maxDbfsBlockSize = 1024 * 1024

// Internal only state for a read handle.
type fileHandleReader struct {
	size   int64
	offset int64
}

func (f *fileHandleReader) errorf(format string, a ...any) error {
	return fmt.Errorf("dbfs read: "+format, a...)
}

func (f *fileHandleReader) error(err error) error {
	if err == nil {
		return nil
	}
	return f.errorf("%w", err)
}

// Internal only state for a write handle.
type fileHandleWriter struct {
	handle int64
}

func (f *fileHandleWriter) errorf(format string, a ...any) error {
	return fmt.Errorf("dbfs write: "+format, a...)
}

func (f *fileHandleWriter) error(err error) error {
	if err == nil {
		return nil
	}
	return f.errorf("%w", err)
}

// Internal only state for a DBFS file handle.
type fileHandle struct {
	ctx  context.Context
	api  *DbfsAPI
	path string

	reader *fileHandleReader
	writer *fileHandleWriter
}

func (h *fileHandle) checkRead() (*fileHandleReader, error) {
	if h.reader != nil {
		return h.reader, nil
	}
	return nil, fmt.Errorf("dbfs: file not open for reading")
}

func (h *fileHandle) checkWrite() (*fileHandleWriter, error) {
	if h.writer != nil {
		return h.writer, nil
	}
	return nil, fmt.Errorf("dbfs: file not open for writing")
}

// Handle defines the interface of the object returned by [DbfsAPI.Open].
type Handle interface {
	io.ReadWriteCloser
	io.WriterTo
	io.ReaderFrom
}

// Implement the [io.Reader] interface.
func (h *fileHandle) Read(p []byte) (int, error) {
	r, err := h.checkRead()
	if err != nil {
		return 0, err
	}

	var ntotal int
	for ntotal < len(p) {
		if r.offset >= r.size {
			return ntotal, io.EOF
		}

		chunk := p[ntotal:]
		if len(chunk) > maxDbfsBlockSize {
			chunk = chunk[:maxDbfsBlockSize]
		}

		res, err := h.api.Read(h.ctx, ReadDbfsRequest{
			Path:   h.path,
			Length: int64(len(chunk)),
			Offset: r.offset,
		})
		if err != nil {
			return ntotal, r.error(err)
		}

		// The guard against offset >= size happens above, so this can only happen
		// if the file is modified or truncated while reading. If this happens,
		// the read contents will likely be corrupted, so we return an error.
		if res.BytesRead == 0 {
			return ntotal, r.errorf("unexpected EOF at offset %d (size %d)", r.offset, r.size)
		}

		nread, err := base64.StdEncoding.Decode(chunk, []byte(res.Data))
		if err != nil {
			return ntotal, r.error(err)
		}

		ntotal += nread
		r.offset += int64(nread)
	}

	return ntotal, nil
}

// Implement the [io.Writer] interface.
func (h *fileHandle) Write(p []byte) (int, error) {
	w, err := h.checkWrite()
	if err != nil {
		return 0, err
	}

	var ntotal int
	for ntotal < len(p) {
		chunk := p[ntotal:]
		if len(chunk) > maxDbfsBlockSize {
			chunk = chunk[:maxDbfsBlockSize]
		}

		err := h.api.AddBlock(h.ctx, AddBlock{
			Data:   base64.StdEncoding.EncodeToString(chunk),
			Handle: w.handle,
		})
		if err != nil {
			return ntotal, w.error(err)
		}

		ntotal += len(chunk)
	}

	return ntotal, nil
}

// Implement the [io.Closer] interface.
func (h *fileHandle) Close() error {
	w, err := h.checkWrite()
	if err != nil {
		return err
	}

	return w.error(h.api.CloseByHandle(h.ctx, w.handle))
}

// Implement the [io.WriterTo] interface.
func (h *fileHandle) WriteTo(w io.Writer) (int64, error) {
	_, err := h.checkRead()
	if err != nil {
		return 0, err
	}

	// Limit types to io.Reader and io.Writer to avoid recursion
	// into WriteTo or ReadFrom functions on underlying types.
	ior := struct{ io.Reader }{h}
	iow := struct{ io.Writer }{w}
	return bufio.NewReaderSize(ior, maxDbfsBlockSize).WriteTo(iow)
}

// Implement the [io.ReaderFrom] interface.
func (h *fileHandle) ReadFrom(r io.Reader) (int64, error) {
	_, err := h.checkWrite()
	if err != nil {
		return 0, err
	}

	// Limit types to io.Reader and io.Writer to avoid recursion
	// into WriteTo or ReadFrom functions on underlying types.
	ior := struct{ io.Reader }{r}
	iow := struct{ io.Writer }{h}
	bw := bufio.NewWriterSize(iow, maxDbfsBlockSize)
	n, err := bw.ReadFrom(ior)
	if err != nil {
		return n, err
	}
	return n, bw.Flush()
}

func (h *fileHandle) openForRead(mode FileMode) error {
	res, err := h.api.GetStatusByPath(h.ctx, h.path)
	if err != nil {
		return err
	}
	if res.IsDir {
		return fmt.Errorf("cannot open directory for reading")
	}
	h.reader = &fileHandleReader{
		size: res.FileSize,
	}
	return nil
}

func (h *fileHandle) openForWrite(mode FileMode) error {
	res, err := h.api.Create(h.ctx, Create{
		Path:      h.path,
		Overwrite: (mode & FileModeOverwrite) != 0,
	})
	if err != nil {
		return err
	}
	h.writer = &fileHandleWriter{
		handle: res.Handle,
	}
	return nil
}

type dbfsAPIUtilities interface {
	// Open opens a remote DBFS file for reading or writing.
	// The returned object implements relevant [io] interfaces for convenient
	// integration with other code that reads or writes bytes.
	//
	// The [io.WriterTo] interface is provided and maximizes throughput for
	// bulk reads by reading data with the DBFS maximum read chunk size of 1MB.
	// Similarly, the [io.ReaderFrom] interface is provided for bulk writing.
	//
	// A file opened for writing must always be closed.
	Open(ctx context.Context, path string, mode FileMode) (Handle, error)

	// ReadFile is identical to [os.ReadFile] but for DBFS.
	ReadFile(ctx context.Context, name string) ([]byte, error)

	// WriteFile is identical to [os.WriteFile] but for DBFS.
	WriteFile(ctx context.Context, name string, data []byte) error

	// RecursiveList traverses the DBFS tree and returns all non-directory
	// objects under the path
	RecursiveList(ctx context.Context, path string) ([]FileInfo, error)
}

// Open opens a remote DBFS file for reading or writing.
// The returned object implements relevant [io] interfaces for convenient
// integration with other code that reads or writes bytes.
//
// The [io.WriterTo] interface is provided and maximizes throughput for
// bulk reads by reading data with the DBFS maximum read chunk size of 1MB.
// Similarly, the [io.ReaderFrom] interface is provided for bulk writing.
//
// A file opened for writing must always be closed.
func (a *DbfsAPI) Open(ctx context.Context, path string, mode FileMode) (Handle, error) {
	h := &fileHandle{
		ctx:  useragent.InContext(ctx, "sdk-feature", "dbfs-handle"),
		api:  a,
		path: path,
	}

	isRead := (mode & FileModeRead) != 0
	isWrite := (mode & FileModeWrite) != 0
	if (isRead && isWrite) || (!isRead && !isWrite) {
		return nil, fmt.Errorf("dbfs open: must specify files.FileModeRead or files.FileModeWrite")
	}

	var err error
	if isRead {
		err = h.openForRead(mode)
	}
	if isWrite {
		err = h.openForWrite(mode)
	}
	if err != nil {
		return nil, fmt.Errorf("dbfs open: %w", err)
	}

	return h, nil
}

// ReadFile is identical to [os.ReadFile] but for DBFS.
func (a *DbfsAPI) ReadFile(ctx context.Context, name string) ([]byte, error) {
	h, err := a.Open(ctx, name, FileModeRead)
	if err != nil {
		return nil, err
	}

	h_ := h.(*fileHandle)
	buf := make([]byte, h_.reader.size)
	_, err = h.Read(buf)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return buf, nil
}

// WriteFile is identical to [os.WriteFile] but for DBFS.
func (a *DbfsAPI) WriteFile(ctx context.Context, name string, data []byte) error {
	h, err := a.Open(ctx, name, FileModeWrite|FileModeOverwrite)
	if err != nil {
		return err
	}

	_, err = h.Write(data)
	cerr := h.Close()
	if err == nil && cerr != nil {
		err = cerr
	}
	return err
}

// RecursiveList traverses the DBFS tree and returns all non-directory
// objects under the path
func (a DbfsAPI) RecursiveList(ctx context.Context, path string) ([]FileInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "recursive-list")
	var results []FileInfo
	queue := []string{path}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		batch, err := a.ListAll(ctx, ListDbfsRequest{
			Path: path,
		})
		if apierr.IsMissing(err) {
			// skip on path deleted during iteration
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("list %s: %w", path, err)
		}
		for _, v := range batch {
			if v.IsDir {
				queue = append(queue, v.Path)
				continue
			}
			results = append(results, v)
		}
	}
	return results, nil
}
