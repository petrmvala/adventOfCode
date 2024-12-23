package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const SecretIter = 2000

func NextSecret(in int) int {
	s1 := ((in * 64) ^ in) % 16777216
	s2 := ((s1 / 32) ^ s1) % 16777216
	out := ((s2 * 2048) ^ s2) % 16777216
	return out
}

func LastSecret(in int) int {
	out := in
	for i := 0; i < SecretIter; i++ {
		out = NextSecret(out)
	}
	return out
}

func SumSecrets(in []int) int {
	out := 0
	for _, s := range in {
		out += LastSecret(s)
	}
	return out
}

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(f), "\n")
	data := []int{}
	for _, l := range lines {
		s, err := strconv.Atoi(l)
		if err != nil {
			if len(l) == 0 {
				continue
			}
			log.Fatal(err)
		}
		data = append(data, s)
	}
	sum := SumSecrets(data)
	fmt.Println(sum)
}
