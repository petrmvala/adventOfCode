package main

import (
	"bufio"
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

func (d Data) SafeReportsDamper() int {
	safe := 0
	for _, report := range d.Reports {
		s, loc := safeReport(report)
		if !s {
			rLeft := make([]int, len(report))
			_ = copy(rLeft, report)
			rLeft = append(rLeft[:loc-1], report[loc:]...)
			sLeft, _ := safeReport(rLeft)
			rRight := make([]int, len(report))
			_ = copy(rRight, report)
			rRight = append(rRight[:loc], report[loc+1:]...)
			sRight, _ := safeReport(rRight)
			rZero := make([]int, len(report))
			_ = copy(rZero, report)
			rZero = report[1:]
			sZero, _ := safeReport(rZero)
			if !(sLeft || sRight || sZero) {
				continue
			}
		}
		safe++
	}
	return safe
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
