package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	numberSeparator = "   "
	fileName        = "data.txt"
)

func main() {
	data, err := NewDataFromFS(os.DirFS("."))
	if err != nil {
		log.Fatal(err)
	}
	if err := data[0].Validate(); err != nil {
		log.Fatal(err)
	}
	dist := data[0].TotalDistance()
	sim := data[0].Similarity()
	fmt.Printf("distance: %d, similarity: %d\n", dist, sim)
}

type Data struct {
	Left  []int
	Right []int
}

func (d Data) Validate() error {
	if len(d.Left) != len(d.Right) {
		return errors.New("invalid data")
	}
	return nil
}

func (d Data) TotalDistance() int {
	sort.Ints(d.Left)
	sort.Ints(d.Right)
	dist := 0
	for i := 0; i < len(d.Left); i++ {
		dist += int(math.Abs(float64(d.Right[i] - d.Left[i])))
	}
	return dist
}

func (d Data) Similarity() int {
	m := map[int]int{}
	for _, r := range d.Right {
		if _, ok := m[r]; !ok {
			m[r] = 0
		}
		m[r]++
	}
	s := 0
	for _, l := range d.Left {
		if v, ok := m[l]; ok {
			s += l * v
		}
	}
	return s
}

func NewDataFromFS(fsys fs.FS) ([]Data, error) {
	dir, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return nil, err
	}
	var data []Data
	for _, f := range dir {
		if f.Name() != fileName {
			continue
		}
		dat, err := getData(fsys, f.Name())
		if err != nil {
			return nil, err
		}
		data = append(data, dat)
	}
	return data, nil
}

func getData(fsys fs.FS, fileName string) (Data, error) {
	dataFile, err := fsys.Open(fileName)
	if err != nil {
		return Data{}, nil
	}
	defer dataFile.Close()
	return newData(dataFile)
}

func newData(data io.Reader) (Data, error) {
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)
	d := Data{}
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), numberSeparator)
		ints := []int{}
		for _, t := range text {
			i, err := strconv.Atoi(t)
			if err != nil {
				return Data{}, nil
			}
			ints = append(ints, i)
		}
		d.Left = append(d.Left, ints[0])
		d.Right = append(d.Right, ints[1])
	}
	return d, nil
}
