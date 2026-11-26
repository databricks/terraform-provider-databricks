package lock

import "github.com/databricks/databricks-sdk-go/qa/lock/internal"

// GitCredentials are unique to the user and workspace.
type GitCredentials struct {
	WorkspaceHost string
	Username      string
}

func (g GitCredentials) GetLockId() string {
	return internal.GenerateBlobName(g)
}

// Some operations require locking the entire workspace.
type Workspace struct {
	WorkspaceId string
}

func (w Workspace) GetLockId() string {
	return internal.GenerateBlobName(w)
}

// LockableImpl is a default implementation of Lockable. If r is a struct, it uses
// reflection to generate a unique lock ID based on the name and fields of the struct.
// If r is a string, it uses the string as the lock ID.
type LockableImpl[R any] struct {
	r R
}

func NewLockable[R any](r R) LockableImpl[R] {
	return LockableImpl[R]{r: r}
}
func (l LockableImpl[R]) GetLockId() string {
	return internal.GenerateBlobName(l.r)
}
