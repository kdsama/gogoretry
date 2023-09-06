package gogoretry

import (
	"errors"
	"testing"
	"time"
)

func RunError(e error) error {
	return e
}

func TestRunErrors(t *testing.T) {
	t.Run("we get an error on all the retries", func(t *testing.T) {
		want := ErrNoResponse
		e := errors.New("Some error")
		ro := New()
		got := ro.Run(func() error {
			return RunError(e)
		})
		if got != want {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})
	t.Run("we shouldnt expect error here ", func(t *testing.T) {

		// e := errors.New("Some error")
		ro := New()
		got := ro.Run(func() error {
			return RunError(nil)
		})
		if got != nil {
			t.Errorf("wanted nil  but got %v", got)
		}
	})

}

func TestRunIntervals(t *testing.T) {
	t.Run("getting an error on all the retres", func(t *testing.T) {

		want := ErrNoResponse
		e := errors.New("Some error")
		ll, err := Custom([]time.Duration{1 * time.Second, 2 * time.Second})
		if err != nil {
			t.Error("Didnt expect error here but got : ", err)
		}
		ro := New(ll)
		got := ro.Run(func() error {
			return RunError(e)
		})
		if got != want {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})
	t.Run("getting an error on all the retres", func(t *testing.T) {

		// e := errors.New("Some error")
		ll, err := Custom([]time.Duration{1 * time.Second, 2 * time.Second})
		if err != nil {
			t.Error("Didnt expect error here but got : ", err)
		}
		ro := New(ll)
		got := ro.Run(func() error {
			return RunError(nil)
		})
		if got != nil {
			t.Errorf("wanted nil  but got %v", got)
		}
	})
}

func TestRetryErrors(t *testing.T) {
	// testSuite

	type testObject struct {
		want    error
		fnInput error
		input   []error
	}
	targetError := errors.New("target error")
	ts := map[string]testObject{
		"Should Return PassedError": {
			want:    targetError,
			fnInput: targetError,
			input:   []error{errors.New("Some other error")},
		},
		"Should return NoResponse": {
			want:    ErrNoResponse,
			fnInput: targetError,
			input:   []error{targetError, errors.New("some random erro1"), errors.New("some random erro2")},
		},
	}
	for name, obj := range ts {
		t.Run(name, func(t *testing.T) {
			ll, err := RetryErrors(obj.input)
			if err != nil {
				t.Error("Expected no error but got ", err)
			}
			r := New(ll)
			got := r.Run(func() error {
				return RunError(obj.fnInput)
			})
			if !errors.Is(obj.want, got) {
				t.Errorf("Expected %v but got %v", obj.want, got)
			}
		})
	}

}

func TestRetyErrors(t *testing.T) {
	t.Run("Should return error ", func(t *testing.T) {
		want := errors.New("Some error")
		ll, _ := RetryErrors([]error{errors.New("Some other error")})
		r := New(ll)
		got := r.Run(func() error {
			return RunError(want)
		})
		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

	})
	t.Run("Should Retry 5 times ", func(t *testing.T) {
		arg := errors.New("Some error")
		want := ErrNoResponse
		ll, err := RetryErrors([]error{arg, errors.New("Some other error2"), errors.New("Some error2")})
		if err != nil {
			t.Error("Expected no error but got ", err)
		}
		r := New(ll)
		got := r.Run(func() error {
			return RunError(arg)
		})
		if !errors.Is(want, got) {
			t.Errorf("Expected %v but got %v", want, got)
		}

	})
}

func TestBadErrors(t *testing.T) {
	// testSuite

	type testObject struct {
		want    error
		fnInput error
		input   []error
	}
	targetError := errors.New("target error")
	ts := map[string]testObject{
		"Should Return PassedError": {
			want:    targetError,
			fnInput: targetError,
			input:   []error{targetError},
		},
		"Should return NoResponse": {
			want:    ErrNoResponse,
			fnInput: errors.New("Random Input"),
			input:   []error{errors.New("some random erro1"), errors.New("some random erro2")},
		},
	}
	for name, obj := range ts {
		t.Run(name, func(t *testing.T) {
			ll, err := BadErrors(obj.input)
			if err != nil {
				t.Error("Expected no error but got ", err)
			}
			r := New(ll)
			got := r.Run(func() error {
				return RunError(obj.fnInput)
			})
			if !errors.Is(obj.want, got) {
				t.Errorf("Expected %v but got %v", obj.want, got)
			}
		})
	}

}
