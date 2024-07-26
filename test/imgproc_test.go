package test

import "testing"

func TestAdd(t *testing.T) {

	got := 1 + 1
	want := 2

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
