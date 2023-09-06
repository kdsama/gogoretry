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

var (
	ErrNoResponse = errors.New("no response from the service")
)

type Action func() error

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

func (r *Retrier) Run(fn Action) error {
	if len(r.intervals) > 0 {
		return r.RunWithIntervals(fn)
	}
	var count int

	var re func(fn Action) error
	re = func(fn Action) error {
		time.Sleep(500 * time.Millisecond)
		if err := fn(); err != nil {

			count++

			if count > r.maxRetries {
				return ErrNoResponse
			}
			return re(fn)
		}
		return nil
	}

	e := re(fn)
	return e
}

func (r *Retrier) RunWithIntervals(fn Action) error {
	var count int

	var re func(fn Action) error
	re = func(fn Action) error {
		time.Sleep(r.intervals[count])
		if err := fn(); err != nil {
			count++
			if count >= r.maxRetries {
				return ErrNoResponse
			}
			return re(fn)
		}
		return nil
	}

	e := re(fn)
	return e
}
