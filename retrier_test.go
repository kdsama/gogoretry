package gogoretry

import (
	"errors"
	"testing"
)

func OverTimeFunc(e error) error {
	return e
}

func TestRunErrors(t *testing.T) {
	t.Run("we get an error on all the retries", func(t *testing.T) {
		want := ErrNoResponse
		e := errors.New("Some error")
		ro := New()
		got := ro.Run(func() error {
			return OverTimeFunc(e)
		})
		if got != want {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})
	t.Run("we shouldnt expect error here ", func(t *testing.T) {

		// e := errors.New("Some error")
		ro := New()
		got := ro.Run(func() error {
			return OverTimeFunc(nil)
		})
		if got != nil {
			t.Errorf("wanted nil  but got %v", got)
		}
	})

}
