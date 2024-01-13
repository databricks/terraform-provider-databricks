package qa

import (
	"context"

	"github.com/databricks/databricks-sdk-go/listing"
)

type SliceIterator[T any] []T

func (s *SliceIterator[T]) HasNext(_ context.Context) bool {
	return len(*s) > 0
}

func (s *SliceIterator[T]) Next(_ context.Context) (T, error) {
	var t T
	if len(*s) == 0 {
		return t, listing.ErrNoMoreItems
	}
	v := (*s)[0]
	*s = (*s)[1:]
	return v, nil
}

var _ listing.Iterator[int] = (*SliceIterator[int])(nil)

type FailingIterator[A any] struct{}

func (f FailingIterator[A]) HasNext(_ context.Context) bool {
	return true
}

func (f FailingIterator[A]) Next(_ context.Context) (A, error) {
	var a A
	return a, ErrImATeapot
}

var _ listing.Iterator[int] = FailingIterator[int]{}
