package gogoretry

import (
	"errors"
	"time"
)

var (
	ErrAtleastOneRetry = errors.New("value should be greater than equal to 1")
	ErrAtleastOneEntry = errors.New("atleast one entry is required")
)

type RetryOpts func(rt *Retrier)

// Set Maximum Number of Retries defaults to 5
func MaxRetries(val int) (RetryOpts, error) {
	if val < 1 {
		return nil, ErrAtleastOneRetry
	}
	return func(rt *Retrier) {
		rt.maxRetries = int(val)
	}, nil
}

func Sleep(t time.Duration) RetryOpts {
	return func(rt *Retrier) {
		rt.sleep = t
	}
}

func Default() RetryOpts {
	return func(rt *Retrier) {
		rt.maxRetries = int(5)
		rt.sleep = 1 * time.Second
	}
}

func Custom(td []time.Duration) (RetryOpts, error) {
	// panic for empty time duration slice
	if len(td) == 0 {
		return nil, ErrAtleastOneEntry
	}
	return func(rt *Retrier) {
		rt.intervals = td
		rt.maxRetries = len(rt.intervals)
	}, nil
}
