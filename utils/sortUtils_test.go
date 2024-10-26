package utils 

import "testing"

func TestUtils(t *testing.T) {
	got := SortString("cab")
	want := "abc"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}