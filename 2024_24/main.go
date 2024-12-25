package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Known map[string]int
type Unknown map[string]Operator

type Operator struct {
	a, b, op string
}

func (o Operator) Operate(known Known) (int, error) {
	var (
		getA, getB int
		ok         bool
	)
	if getA, ok = known[o.a]; !ok {
		return -1, errors.New("not found")
	}
	if getB, ok = known[o.b]; !ok {
		return -1, errors.New("not found")
	}
	switch o.op {
	case "AND":
		return getA & getB, nil
	case "OR":
		return getA | getB, nil
	case "XOR":
		return getA ^ getB, nil
	default:
		log.Fatalln("this is weird")
	}
	return -1, errors.New("never happens")
}

func (o Operator) String() string {
	return fmt.Sprintf("%s{a:%s, b:%s}", o.op, o.a, o.b)
}

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(GetZet(string(f)))
}

func GetZet(input string) int {
	known := Known{}
	unknown := Unknown{}

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, ":") {
			s := strings.Split(line, ": ")
			val, err := strconv.ParseInt(s[1], 2, 0)
			if err != nil {
				log.Fatal(err)
			}
			known[s[0]] = int(val)
			continue
		}
		expression := strings.Split(line, " -> ")
		out := expression[1]
		in := strings.Split(expression[0], " ")
		unknown[out] = Operator{a: in[0], b: in[2], op: in[1]}
	}

	for len(unknown) > 0 {
		for k, v := range unknown {
			out, err := v.Operate(known)
			if err != nil {
				continue
			}
			known[k] = out
			delete(unknown, k)
			break
		}
	}

	zet := strings.Repeat("0", 50)
	insertBit := func(bit string, address int) string {
		last := len(zet) - 1
		return zet[:last-address] + bit + zet[last-address+1:]
	}

	for k, v := range known {
		if !strings.HasPrefix(k, "z") {
			continue
		}
		z, err := strconv.ParseInt(k[1:], 10, 0)
		if err != nil {
			log.Fatal(err)
		}
		zet = insertBit(strconv.Itoa(v), int(z))
	}

	res, err := strconv.ParseInt(zet, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(res)
}
