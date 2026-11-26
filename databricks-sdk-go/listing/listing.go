package listing

import (
	"context"
	"errors"
)

// Iterator[T] is an iterator over items of type T. It is similar to
// java.util.Iterator. It is not thread-safe.
type Iterator[T any] interface {
	// HasNext returns true if there are more items to iterate over. HasNext
	// will also return true if the iterator needs to fetch the next page of
	// items from the underlying source and the request fails, even if there
	// are no more items to iterate over. In this case, Next will return the
	// error.
	HasNext(context.Context) bool

	// Next returns the next item in the iterator. If there are no more items
	// to iterate over, it returns ErrNoMoreItems. If there was an error
	// fetching the next page of items, it returns that error. Once Next
	// returns ErrNoMoreItems or an error, it will continue to return that
	// error.
	Next(context.Context) (T, error)
}

// ToSlice returns all items from the iterator as a slice. If there was an
// error fetching items at any time, it returns that error.
func ToSlice[T any](ctx context.Context, it Iterator[T]) ([]T, error) {
	var items []T
	for it.HasNext(ctx) {
		item, err := it.Next(ctx)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// ToSliceN returns the first N items from the iterator as a slice. If there
// was an error fetching items at any time, it returns that error. If n == 0,
// it returns all items.
func ToSliceN[T any, Limit ~int | ~int64](ctx context.Context, it Iterator[T], n Limit) ([]T, error) {
	if n == 0 {
		return ToSlice(ctx, it)
	}
	var items []T
	for it.HasNext(ctx) && Limit(len(items)) < n {
		item, err := it.Next(ctx)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// Use struct{} for Req for iterators that don't need any request structure.
type PaginatingIterator[Req, Resp, T any] struct {
	// nextReq is the request to be used to fetch the next page. If nil, then
	// there is no next page to fetch.
	nextReq *Req

	// getNextPage fetches the next page of items, returning the deserialized
	// response and error.
	getNextPage func(context.Context, Req) (Resp, error)

	// getItems selects the items being iterated over from the response.
	getItems func(Resp) []T

	// getNextReq is used to get the next request to be used in the next page.
	// If nil, then there is no next page to fetch.
	getNextReq func(Resp) *Req

	// currentPage is the current page of items.
	currentPage []T

	// currentPageIdx is the index of the next item from currentPage to return.
	currentPageIdx int

	// lastErr is the last error returned by getNextPage.
	lastErr error
}

var ErrNoMoreItems = errors.New("no more items")

// Returns a new iterator. The iterator will fetch the next page of items
// lazily, when needed. nextReq is the request to be used to fetch the initial
// page. If nil, then no page will be fetched. getNextPage fetches the next
// page of items, returning the deserialized response and error. getItems
// selects the items being iterated over from the response. getNextReq is used
// to get the next request to be used in the next page. If the returned value
// is nil, then there is no next page to fetch.
func NewIterator[Req, Resp, T any](
	nextReq *Req,
	getNextPage func(context.Context, Req) (Resp, error),
	getItems func(Resp) []T,
	getNextReq func(Resp) *Req,
) *PaginatingIterator[Req, Resp, T] {
	return &PaginatingIterator[Req, Resp, T]{
		nextReq:     nextReq,
		getNextPage: getNextPage,
		getItems:    getItems,
		getNextReq:  getNextReq,
	}
}

func (i *PaginatingIterator[Req, Resp, T]) loadNextPageIfNeeded(ctx context.Context) error {
	if i.currentPageIdx < len(i.currentPage) {
		return nil
	}
	if i.nextReq == nil {
		i.currentPage = nil
		i.currentPageIdx = 0
		return nil
	}
	if i.lastErr != nil {
		return i.lastErr
	}

	// Keep loading pages while we have a next request and the current page is
	// empty.
	i.currentPage = nil
	i.currentPageIdx = 0
	// Endpoints using token-based pagination may return an empty page with a
	// next token. We need to keep fetching pages until we get a non-empty
	// page or there are no more pages.
	for i.nextReq != nil && len(i.currentPage) == 0 {
		resp, err := i.getNextPage(ctx, *i.nextReq)
		i.lastErr = err
		if err != nil {
			return err
		}
		if i.getNextReq != nil {
			i.nextReq = i.getNextReq(resp)
		} else {
			i.nextReq = nil
		}
		i.currentPage = i.getItems(resp)
	}
	return nil
}

func (i *PaginatingIterator[Req, Resp, T]) Next(ctx context.Context) (T, error) {
	var t T
	err := i.loadNextPageIfNeeded(ctx)
	if err != nil {
		return t, err
	}
	if i.currentPageIdx >= len(i.currentPage) {
		return t, ErrNoMoreItems
	}
	item := i.currentPage[i.currentPageIdx]
	i.currentPageIdx++
	return item, nil
}

func (i *PaginatingIterator[Req, Resp, T]) HasNext(ctx context.Context) bool {
	err := i.loadNextPageIfNeeded(ctx)
	// As described in the documentation for HasNext, if there was an error
	// fetching the next page, we still return true to allow the user to handle
	// the error in Next.
	if err != nil {
		return true
	}
	return i.currentPageIdx < len(i.currentPage)
}

type DeduplicatingIterator[T any, Id comparable] struct {
	it      Iterator[T]
	getId   func(T) Id
	seen    map[Id]struct{}
	current *T
}

func NewDedupeIterator[T any, Id comparable](it Iterator[T], getId func(T) Id) *DeduplicatingIterator[T, Id] {
	return &DeduplicatingIterator[T, Id]{
		it:    it,
		getId: getId,
		seen:  make(map[Id]struct{}),
	}
}

func (i *DeduplicatingIterator[T, Id]) Next(ctx context.Context) (T, error) {
	if i.current != nil {
		t := *i.current
		i.current = nil
		return t, nil
	}
	for {
		t, err := i.it.Next(ctx)
		if err != nil {
			return t, err
		}
		id := i.getId(t)
		if _, ok := i.seen[id]; !ok {
			i.seen[id] = struct{}{}
			return t, nil
		}
	}
}

func (i *DeduplicatingIterator[T, Id]) HasNext(ctx context.Context) bool {
	if i.current != nil {
		return true
	}
	// To compute HasNext in DeduplicatingIterator, we need to actually fetch
	// the next item from the underlying iterator and compare it to seen items.
	// However, the retrieved item cannot be discarded, as it needs to be
	// returned by the next call to Next. So we store the item in current and
	// return it in the next call to Next, after which current is set to nil.
	for {
		t, err := i.it.Next(ctx)
		if errors.Is(err, ErrNoMoreItems) {
			return false
		}
		if err != nil {
			return true
		}
		id := i.getId(t)
		if _, ok := i.seen[id]; !ok {
			i.current = &t
			i.seen[id] = struct{}{}
			return true
		}
	}
}

// A simple iterator over a slice.
type SliceIterator[T any] []T

func (s *SliceIterator[T]) HasNext(_ context.Context) bool {
	return len(*s) > 0
}

func (s *SliceIterator[T]) Next(_ context.Context) (T, error) {
	var t T
	if len(*s) == 0 {
		return t, ErrNoMoreItems
	}
	v := (*s)[0]
	*s = (*s)[1:]
	return v, nil
}

type limitIterator[T any] struct {
	it        Iterator[T]
	remaining int
}

// NewLimitIterator wraps an iterator and yields at most n items. If n <= 0,
// the inner iterator is returned unchanged (no limit is applied). Note that
// this differs from ToSliceN, where n == 0 means unlimited but negative n
// yields zero items.
func NewLimitIterator[T any](iter Iterator[T], n int) Iterator[T] {
	if n <= 0 {
		return iter
	}
	return &limitIterator[T]{
		it:        iter,
		remaining: n,
	}
}

func (i *limitIterator[T]) HasNext(ctx context.Context) bool {
	if i.remaining <= 0 {
		return false
	}
	return i.it.HasNext(ctx)
}

func (i *limitIterator[T]) Next(ctx context.Context) (T, error) {
	var t T
	if i.remaining <= 0 {
		return t, ErrNoMoreItems
	}
	t, err := i.it.Next(ctx)
	if err != nil {
		return t, err
	}
	i.remaining--
	return t, nil
}
