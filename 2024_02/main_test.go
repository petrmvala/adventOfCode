package main

import (
	"reflect"
	"testing"
	"testing/fstest"
)

const sampleData = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

var sample = Data{
	Reports: [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	},
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

func TestSafeReports(t *testing.T) {
	data := sample
	got := data.SafeReports()
	want := 2
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSafeReportsDamper(t *testing.T) {
	data := sample
	got := data.SafeReportsDamper()
	want := 4
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
