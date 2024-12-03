package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const fileName = "data.txt"

func main() {
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	input := string(f)
	m := Multiply(input)
	e := Enabled(input)
	fmt.Println("multiply: ", m)
	fmt.Println("enabled: ", e)
}

func Multiply(data string) int {
	r, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		log.Fatal(err)
	}
	multiplys := r.FindAllStringSubmatch(data, -1)
	mul := 0
	for _, m := range multiplys {
		ints := toInts(m)
		mul += ints[0] * ints[1]
	}
	return mul
}

func Enabled(data string) int {
	r, err := regexp.Compile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		log.Fatal(err)
	}
	matches := r.FindAllStringSubmatch(data, -1)
	enabled := true
	mul := 0
	for _, m := range matches {
		if m[0] == "do()" {
			enabled = true
			continue
		} else if m[0] == "don't()" {
			enabled = false
			continue
		}
		if enabled {
			ints := toInts(m)
			mul += ints[0] * ints[1]
		}
	}
	return mul
}

func toInts(s []string) []int {
	l, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}
	r, err := strconv.Atoi(s[2])
	if err != nil {
		log.Fatal(err)
	}
	return []int{l, r}
}
