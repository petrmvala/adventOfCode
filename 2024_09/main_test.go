package main

import (
	"testing"
)

const (
	smallInput = "2333133121414131402"
)

func TestSumDefrag(t *testing.T) {
	got := SumDefrag(smallInput)
	want := 1928
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
