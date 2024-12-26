package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	keyLine  = "....."
	lockLine = "#####"
)

type Schematic []int

func GetSchematics(input string) ([]Schematic, []Schematic) {
	keys := []Schematic{}
	locks := []Schematic{}
	for _, schematic := range strings.Split(input, "\n\n") { // range through schematics
		rows := strings.Split(schematic, "\n") // this is a whole schematic including padding rows (rows[0] and rows[6])
		s := Schematic{0, 0, 0, 0, 0}
		for _, pinRow := range rows[1:6] { // range through the rows with pins only
			for pin, val := range strings.Split(pinRow, "") {
				if val == "#" {
					s[pin]++
				}
			}
		}
		if rows[0] == lockLine && rows[6] == keyLine { // locks have top row filled with # and bottom with .
			locks = append(locks, s)
		} else {
			keys = append(keys, s)
		}
	}
	return keys, locks
}

func GetPairs(input string) int {
	keys, locks := GetSchematics(input)
	sum := 0
	for _, key := range keys {
		for _, lock := range locks {
			for i := 0; i < len(key); i++ {
				if key[i]+lock[i] > 5 {
					break
				}
				if i == len(key)-1 {
					sum++
				}
			}
		}
	}
	return sum
}

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	p := GetPairs(string(f))
	fmt.Println(p)
}
