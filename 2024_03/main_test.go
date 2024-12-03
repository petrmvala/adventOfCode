package main

import (
	"testing"
)

const sampleData = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
const sampleData2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestMultiply(t *testing.T) {
	got := Multiply(sampleData)
	want := 161
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestEnabled(t *testing.T) {
	got := Enabled(sampleData2)
	want := 48
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
