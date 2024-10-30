package utils

import "testing"

func TestSortUtils(t *testing.T) {
	got := SortString("cab")
	want := "abc"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMax(t *testing.T) {
	got := Max([]int{1, 2, 3, 4, 5})
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3, 4, 5})
	want := 15

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestMaxNum(t *testing.T) {
	got := MaxNum(2, 3)
	want := 3

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestMinNum(t *testing.T) {
	got := MinNum(2, 3)
	want := 2

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
