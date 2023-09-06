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

// Set Maximum Number of Retries. Returns error for < 1
func MaxRetries(val int) (RetryOpts, error) {
	if val < 1 {
		return nil, ErrAtleastOneRetry
	}
	return func(rt *Retrier) {
		rt.maxRetries = int(val)
	}, nil
}

// Sets sleep Time
func Sleep(t time.Duration) RetryOpts {
	return func(rt *Retrier) {
		rt.sleep = t
	}
}

// Default configuration
func Default() RetryOpts {
	return func(rt *Retrier) {
		rt.maxRetries = int(5)
		rt.sleep = 1 * time.Second
	}
}

// Custom time intervals. Retrier will automatically set maxRetry
// according to the length of the input. Returns error for empty slice
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

func BadErrors(arr []error) (RetryOpts, error) {

	return func(rt *Retrier) {
		// set retryErrorFlag to false
		rt.re = false
		rt.be = true
		for _, e := range arr {
			rt.badErrors[e] = true
		}

	}, nil

}

func RetryErrors(arr []error) (RetryOpts, error) {
	return func(rt *Retrier) {
		// set badErrorFlag to false
		rt.be = false
		rt.re = true
		for _, e := range arr {
			rt.retryErrors[e] = true
		}

	}, nil
}
