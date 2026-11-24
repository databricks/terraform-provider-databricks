package poll

import "time"

type PollFunc[R any] func(time.Duration, func(*R)) (*R, error)

func Simple[R any](r R) PollFunc[R] {
	return func(_ time.Duration, _ func(*R)) (*R, error) {
		return &r, nil
	}
}

func SimpleError[R any](err error) PollFunc[R] {
	return func(_ time.Duration, _ func(*R)) (*R, error) {
		return nil, err
	}
}
