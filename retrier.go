package gogoretry

import (
	"errors"
	"time"
)

// So we need to make it cooler than before
// We can have a normal one, Max number of retries, and a multiplier
// By default we can set this to 5 and 1
// Also what we can have is Custom. THis is where we gonna give
// so functional Pattern is basically what we gonna pass
//

type Retrier struct {
	// Sleep time between each retry . Defaults to 1 second
	sleep time.Duration
	// Maximum number of retries
	maxRetries int
	// Custom Intervals allow
	intervals []time.Duration
}

func New(opts ...RetryOpts) *Retrier {

	r := &Retrier{}
	// Load default settings
	Default()(r)
	for _, op := range opts {
		op(r)
	}
	return r
}

type opts func() error

func (r *Retrier) Run(fn opts) error {
	if len(r.intervals) > 0 {
		return r.RunWithIntervals(fn)
	}
	var count int

	var re func(fn opts) error
	re = func(fn opts) error {
		time.Sleep(500 * time.Millisecond)
		if err := fn(); err != nil {

			count++

			if count > r.maxRetries {
				return errors.New("Oh how damn ")
			}
			re(fn)
		}
		return nil
	}

	e := re(fn)
	return e
}

func (r *Retrier) RunWithIntervals(fn opts) error {
	var count int

	var re func(fn opts) error
	re = func(fn opts) error {
		time.Sleep(r.intervals[count])
		if err := fn(); err != nil {
			count++
			if count > r.maxRetries {
				return errors.New("limit exceeded")
			}
			re(fn)
		}
		return nil
	}

	e := re(fn)
	return e
}
