package gogoretry

import (
	"testing"
	"time"
)

func TestMaxRetriesError(t *testing.T) {
	_, got := MaxRetries(-1)
	want := ErrAtleastOneRetry
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}

}
func TestMaxRetries(t *testing.T) {
	r, got := MaxRetries(1)

	if got != nil {
		t.Errorf("Wanted nil  but got %v", got)
	}

	retryObj := New(r)
	if retryObj.maxRetries != 1 {
		t.Errorf("Wanted 1 but got %v", retryObj.maxRetries)
	}

}

func TestCustomErrors(t *testing.T) {
	_, got := Custom([]time.Duration{})
	if got != ErrAtleastOneEntry {
		t.Errorf("wanted %v but got %v", ErrAtleastOneEntry, got)
	}
}

func TestCustom(t *testing.T) {
	arr := []time.Duration{1 * time.Second, 5 * time.Second}
	r, err := Custom(arr)
	if err != nil {
		t.Errorf("wanted nil but got %v", err)
	}
	retryObj := New(r)
	if len(retryObj.intervals) != len(arr) {
		t.Error("Mismatch in length")
	}
	for i := range retryObj.intervals {
		if arr[i] != retryObj.intervals[i] {
			t.Errorf("%d th element does not match", i)
		}
	}

}
