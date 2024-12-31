package main

import (
	"reflect"
	"testing"
)

const smallInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

var dataInput = map[int][]int{
	190:    {10, 19},
	3267:   {81, 40, 27},
	83:     {17, 5},
	156:    {15, 6},
	7290:   {6, 8, 6, 15},
	161011: {16, 10, 13},
	192:    {17, 8, 14},
	21037:  {9, 7, 18, 13},
	292:    {11, 6, 16, 20},
}

func TestData(t *testing.T) {
	got := Data(smallInput)
	want := dataInput
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	got := Sum(Data(smallInput))
	want := 3749
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
