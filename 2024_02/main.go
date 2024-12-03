package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	fileName        = "data.txt"
	numberSeparator = " "
)

var ErrIndexOutOfBounds = errors.New("index out of bounds")

func main() {
	data, err := NewDataFromFS(os.DirFS("."))
	if err != nil {
		log.Fatal(err)
	}
	rep := data[0].SafeReports()
	dam := data[0].SafeReportsDamper()
	fmt.Printf("SafeReports: %d, SafeReportsDamper: %d\n", rep, dam)
}

type Data struct {
	Reports [][]int
}

func safeReport(report []int) (bool, int) {
	direction := 1
	previous := report[0]
	if report[0] > report[1] {
		direction = -1
	}
	for i, current := range report {
		if i == 0 {
			continue
		}
		q := direction * (current - previous)
		if q < 1 || q > 3 {
			return false, i
		}
		previous = current
	}
	return true, 0
}

func (d Data) SafeReports() int {
	safe := 0
	for _, report := range d.Reports {
		if s, _ := safeReport(report); !s {
			continue
		}
		safe++
	}
	return safe
}

func damper(report []int) bool {
	s, loc := safeReport(report)
	if !s {
		for _, idx := range []int{0, loc - 1, loc} {
			r, err := removeIndex(report, idx)
			if err != nil {
				log.Fatal(err)
			}
			s, _ = safeReport(r)
			if s {
				return true
			}
		}
		return false
	}
	return true
}

func (d Data) SafeReportsDamper() int {
	safe := 0
	for _, report := range d.Reports {
		if damper(report) {
			safe++
		}
	}
	return safe
}

// This did bite, check answer on SO
// https://stackoverflow.com/a/57213476/3027202
func removeIndex(slice []int, index int) ([]int, error) {
	if index < 0 || len(slice) <= index {
		return nil, ErrIndexOutOfBounds
	}
	ret := make([]int, 0)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...), nil
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
		d.Reports = append(d.Reports, ints)
	}
	return d, nil
}
