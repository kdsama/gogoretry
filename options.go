package gogoretry

import (
	"errors"
	"time"
)

var (
	//ErrAtleastOne is an error for when values passed is less than 1
	ErrAtleastOne = errors.New("value should be greater than equal to 1")

	//ErrAtleastOneEntry is an error for when atleast one entry,
	// usually in a slice, is required.
	ErrAtleastOneEntry = errors.New("atleast one entry is required")
)

// RetryOpts is used to set configuration of Retrier
type RetryOpts func(rt *Retrier)

// MaxRetries sets Maximum Number of Retries. Returns error for < 1
func MaxRetries(val int) (RetryOpts, error) {
	if val < 1 {
		return nil, ErrAtleastOne
	}
	return func(rt *Retrier) {
		rt.maxRetries = int(val)
	}, nil
}

// Sleep sets sleep Time
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

// Exponential backoffs for
func Exponential(t time.Duration, multiplier int, maxRetries int) RetryOpts {
	var intervals = make([]time.Duration, maxRetries)
	intervals[0] = t
	for i := 1; i < maxRetries; i++ {
		intervals[i] = intervals[i-1] * time.Duration(multiplier)
	}
	return func(rt *Retrier) {
		rt.intervals = intervals
	}
}

// BadErrors sets errorSlice in retrier. Whenever these errors
// come up, Retrier returns instead of retrying
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

// RetryErrors sets an errorSlice in retrier. Whenever these errors
// dont come up, Retrier returns instead of retrying.
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
