package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(string(f), "\n")
	fmt.Println("diskmap")
	dm := DiskMap(data[0])
	fmt.Println("defrag")
	df := Defrag(dm)
	fmt.Println(df)
	fmt.Println("Checksum:", Checksum(df))
}

func DiskMap(input string) string {
	m := ""
	for idx, val := range strings.Split(input, "") {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		n := "."
		if idx%2 == 0 { // current value is a file block
			n = strconv.Itoa(idx / 2)
		}
		m += strings.Repeat(n, v)
	}
	return m
}

func Defrag(input string) string {
	// fmt.Println("input:", input)
	d := strings.Split(input, "")
	stop := len(d) - 1
	for i := 0; i < stop; i++ {
		// fmt.Println("forward index:", i, "stop:", stop, "d[i]:", d[i])
		if d[i] != "." {
			continue
		}
		for d[stop] == "." {
			stop--
			// fmt.Println("backward index:", stop, "d[stop]:", d[stop])
		}
		tmp := d[stop]
		d = slices.Concat(d[:i], []string{tmp}, d[i+1:stop], []string{"."}, d[stop+1:])
		stop--
	}
	return strings.Join(d, "")
}

func Checksum(input string) int {
	d := strings.Split(input, "")
	sum := 0
	for i, n := range d {
		if n == "." {
			break
		}
		val, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		sum += i * val
	}
	return sum
}
