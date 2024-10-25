package main 

import "testing"

func TestSortUtils(t *testing.T) {
	got := sortString("cab")
	want := "abc"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}