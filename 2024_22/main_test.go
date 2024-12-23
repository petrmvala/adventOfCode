package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestNextSecret(t *testing.T) {
	want := 15887950
	got := NextSecret(123)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestNextSecretSequence(t *testing.T) {
	want := []int{15887950, 16495136, 527345, 704524, 1553684, 12683156, 11100544, 12249484, 7753432, 5908254}
	got := []int{}
	g := 123
	for i := 0; i < len(want); i++ {
		g = NextSecret(g)
		got = append(got, g)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetLastSecretNumber(t *testing.T) {
	data := []struct {
		in  int
		out int
	}{
		{1, 8685429},
		{10, 4700978},
		{100, 15273692},
		{2024, 8667524},
	}
	for _, tt := range data {
		t.Run(fmt.Sprintf("input: %s", strconv.Itoa(tt.in)), func(t *testing.T) {
			got := LastSecret(tt.in)
			if got != tt.out {
				t.Errorf("got %d, want %d", got, tt.out)
			}
		})
	}
}

func TestGetSecretSum(t *testing.T) {
	got := SumSecrets([]int{1, 10, 100, 2024})
	want := 37327623
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
