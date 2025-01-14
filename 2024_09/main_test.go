package main

import (
	"testing"
)

const (
	smallInput   = "2333133121414131402"
	smallMap     = "00...111...2...333.44.5555.6666.777.888899"
	defragmented = "0099811188827773336446555566.............."
)

func TestDiskMap(t *testing.T) {
	got := DiskMap(smallInput)
	want := smallMap
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestDefrag(t *testing.T) {
	got := Defrag(smallMap)
	want := defragmented
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestChecksum(t *testing.T) {
	got := Checksum(defragmented)
	want := 1928
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
