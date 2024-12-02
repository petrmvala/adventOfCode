package main

import (
	"reflect"
	"testing"
	"testing/fstest"
)

const (
	sampleData = `3   4
4   3
2   5
1   3
3   9
3   3
`
)

var sample = Data{
	Left:  []int{3, 4, 2, 1, 3, 3},
	Right: []int{4, 3, 5, 3, 9, 3},
}

func TestNewDataFromFS(t *testing.T) {
	fs := fstest.MapFS{fileName: {Data: []byte(sampleData)}}
	data, err := NewDataFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) != len(fs) {
		t.Errorf("got %d data, want %d data", len(data), len(fs))
	}
	got := data[0]
	want := sample
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestTotalDistance(t *testing.T) {
	data := sample
	got := data.TotalDistance()
	want := 11
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimilarity(t *testing.T) {
	data := sample
	got := data.Similarity()
	want := 31
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
