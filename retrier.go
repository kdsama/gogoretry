package gogoretry

import (
	"errors"
	"time"
)

var (
	//ErrNoResponse is the error returned when we have retried
	// max number of times.
	ErrNoResponse = errors.New("no response from the service")
)

// Action encapsulates the user method.
// For user to implement it, they would have to wrap their function inside Action.
type Action func() error

// Retrier is the original struct will all the settings related to the
// Gogoretrier package.
type Retrier struct {
	// Sleep time between each retry . Defaults to 1 second
	sleep time.Duration

	// Maximum number of retries
	maxRetries int

	// Custom Intervals allow
	intervals []time.Duration

	// be sets badErrors true or false. Will be automatically disabled
	// if the other one is enabled
	be bool

	// If any error from this list pops up, there  wont be any retry
	badErrors map[error]bool

	// be sets retryErrors true or false. Will be automatically disabled
	// if the other one is enabled
	re bool

	// If list is shared, any error in retryErrors will be retried.
	//Anyother error will lead to termination
	retryErrors map[error]bool
}

// New initiates a retrier.
func New(opts ...RetryOpts) *Retrier {

	r := &Retrier{
		badErrors:   map[error]bool{},
		retryErrors: map[error]bool{},
	}
	// Load default settings
	Default()(r)
	for _, op := range opts {
		op(r)
	}
	return r
}

// Run runs the user method wrapped inside Action
// with a set number of retries according to the configuration
func (r *Retrier) Run(fn Action) error {

	if len(r.intervals) > 0 {
		return r.RunWithIntervals(fn)
	}
	var (
		count       int
		badErrors   = r.badErrors
		be          = r.be
		maxRetries  = r.maxRetries
		re          = r.re
		retryErrors = r.retryErrors
		sleep       = r.sleep
	)

	var rn func(fn Action) error
	rn = func(fn Action) error {

		if err := fn(); err != nil {
			if be {
				if _, ok := badErrors[err]; ok {
					return err
				}
			}

			if re {
				if _, ok := retryErrors[err]; !ok {
					return err
				}
			}
			count++

			if count > maxRetries {
				return ErrNoResponse
			}
			time.Sleep(sleep)
			return rn(fn)
		}
		return nil
	}

	e := rn(fn)
	return e
}

// RunWithIntervals is similar to Run. Difference is that we have a slice
// of time durations corresponding to each retry here, instead of maxRetries
// and constant sleep gap.
func (r *Retrier) RunWithIntervals(fn Action) error {
	var (
		count       int
		badErrors   = r.badErrors
		be          = r.be
		maxRetries  = r.maxRetries
		re          = r.re
		retryErrors = r.retryErrors
		intervals   = r.intervals
	)

	var rn func(fn Action) error
	rn = func(fn Action) error {

		if err := fn(); err != nil {
			if be {
				if _, ok := badErrors[err]; ok {
					return err
				}
			}

			if re {
				if _, ok := retryErrors[err]; !ok {
					return err
				}
			}
			count++
			if count >= maxRetries {
				return ErrNoResponse
			}
			time.Sleep(intervals[count])
			return rn(fn)
		}
		return nil
	}

	e := rn(fn)
	return e
}
