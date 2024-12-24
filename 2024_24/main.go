package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Operator interface {
	Operate(known map[string]bool) (bool, error)
}

type And struct {
	a, b string
}

func (a And) String() string {
	return fmt.Sprintf("And{a:%s, b:%s}", a.a, a.b)
}
func (a And) Operate(known map[string]bool) (bool, error) {
	getA, ok := known[a.a]
	if !ok {
		return false, errors.New("not found")
	}
	getB, ok := known[a.b]
	if !ok {
		return false, errors.New("not found")
	}
	return getA && getB, nil
}

type Or struct {
	a, b string
}

func (o Or) String() string {
	return fmt.Sprintf("Or{a:%s, b:%s}", o.a, o.b)
}
func (o Or) Operate(known map[string]bool) (bool, error) {
	getA, ok := known[o.a]
	if !ok {
		return false, errors.New("not found")
	}
	getB, ok := known[o.b]
	if !ok {
		return false, errors.New("not found")
	}
	return getA || getB, nil
}

type Xor struct {
	a, b string
}

func (x Xor) String() string {
	return fmt.Sprintf("Xor{a:%s, b:%s}", x.a, x.b)
}
func (x Xor) Operate(known map[string]bool) (bool, error) {
	getA, ok := known[x.a]
	if !ok {
		return false, errors.New("not found")
	}
	getB, ok := known[x.b]
	if !ok {
		return false, errors.New("not found")
	}
	return getA != getB, nil
}

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(GetZet(string(f)))
}

func GetZet(input string) int {
	known := map[string]bool{}

	parts := strings.Split(input, "\n\n")
	for _, line := range strings.Split(parts[0], "\n") {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, ":") {
			s := strings.Split(line, ": ")
			val, err := strconv.ParseBool(s[1])
			if err != nil {
				log.Fatal(err)
			}
			known[s[0]] = val
		}
	}

	// fmt.Printf("known: %+v\n", known)

	unknown := map[string]Operator{}
	for _, line := range strings.Split(parts[1], "\n") {
		if len(line) == 0 {
			continue
		}
		expression := strings.Split(line, " -> ")
		out := expression[1]
		in := strings.Split(expression[0], " ")
		switch in[1] {
		case "AND":
			unknown[out] = And{a: in[0], b: in[2]}
		case "OR":
			unknown[out] = Or{a: in[0], b: in[2]}
		case "XOR":
			unknown[out] = Xor{a: in[0], b: in[2]}
		default:
			log.Fatalln("this is weird")
		}
	}

	// fmt.Printf("unknown: %+v\n", unknown)

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

	// fmt.Printf("known: %+v\n", known)
	// fmt.Printf("unknown: %+v\n", unknown)

	type zetBit struct {
		bit string
		val bool
	}
	zetBits := []zetBit{}
	for k, v := range known {
		if !strings.HasPrefix(k, "z") {
			continue
		}
		zetBits = append(zetBits, zetBit{bit: k, val: v})
	}
	sort.Slice(zetBits, func(i int, j int) bool {
		// reverse sort by string
		return zetBits[i].bit > zetBits[j].bit
	})

	// fmt.Printf("sorted zetBits: %+v\n", zetBits)

	var b strings.Builder
	for _, bit := range zetBits {
		if bit.val {
			fmt.Fprintf(&b, "%b", 1)
		} else {
			fmt.Fprintf(&b, "%b", 0)
		}
	}

	binstr := b.String()
	// fmt.Printf("result: %s", binstr)

	res, err := strconv.ParseInt(binstr, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(res)

}
