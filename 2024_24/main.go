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

type Inputs map[string]int
type Operations map[string]Operator

type Operator struct {
	a, b, op string
}

func (o Operator) Operate(known Inputs) (int, error) {
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
	inputs, ops := LoadInput(string(f))

	fmt.Println("Analysis:", AnalyzeAdders(ops))

	fmt.Println(GetZet(inputs, ops))
}

func LoadInput(input string) (Inputs, Operations) {
	inputs := Inputs{}
	operations := Operations{}

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
			inputs[s[0]] = int(val)
			continue
		}
		expression := strings.Split(line, " -> ")
		out := expression[1]
		in := strings.Split(expression[0], " ")
		operations[out] = Operator{a: in[0], b: in[2], op: in[1]}
	}

	return inputs, operations
}

func GetZet(inputs Inputs, operations Operations) int {
	for len(operations) > 0 {
		for k, v := range operations {
			out, err := v.Operate(inputs)
			if err != nil {
				continue
			}
			inputs[k] = out
			delete(operations, k)
			break
		}
	}

	zet := strings.Repeat("0", 50)
	insertBit := func(bit string, address int) string {
		last := len(zet) - 1
		return zet[:last-address] + bit + zet[last-address+1:]
	}

	for k, v := range inputs {
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

func AnalyzeAdders(ops Operations) string {
	output := []string{}
	fixes := Operations{
		// scw XOR gnj got shj, want z07
		"z07": Operator{a: "scw", b: "gnj", op: "XOR"},
		"shj": Operator{a: "scw", b: "gnj", op: "AND"},
		// wkb XOR ngf got , want z16
		"wkb": Operator{a: "x16", b: "y16", op: "AND"},
		"tpk": Operator{a: "x16", b: "y16", op: "XOR"},
		// pvr XOR gvb got pfn, want z23
		// Finished, CarryOut: z23
		"z23": Operator{a: "pvr", b: "gvb", op: "XOR"},
		"pfn": Operator{a: "jwb", b: "hjp", op: "OR"},
		// vpt XOR hdg got kcd, want z27
		// gtc OR z27 not found
		"z27": Operator{a: "vpt", b: "hdg", op: "XOR"},
		"kcd": Operator{a: "y27", b: "x27", op: "AND"},
	}

	for out, op := range fixes {
		ops[out] = op
		output = append(output, out)
	}
	// z00 is a special case
	// x00 XOR y00 -> z00
	// x00 AND y00 -> dhb == CarryOut
	// start from 1 and initialize carry with dhb
	carry := "dhb"
	for i := 1; i < 45; i++ {
		s := strconv.Itoa(i)
		if i < 10 {
			s = "0" + s
		}
		a := Adder{
			A:       "x" + s,
			B:       "y" + s,
			Sum:     "z" + s,
			CarryIn: carry,
		}
		fmt.Printf("Processing Adder %+v\n", a)
		a.Validate(ops)
		if carry != a.CarryIn {
			fmt.Printf("Adder %+v expected CarryIn = %s\n", a, carry)
		}
		carry = a.CarryOut
		fmt.Printf("Finished, CarryOut: %s\n", a.CarryOut)
	}

	sort.Strings(output)
	return strings.Join(output, ",")
}

type Adder struct {
	A        string
	B        string
	CarryIn  string
	CarryOut string
	Sum      string
}

func (a *Adder) Validate(ops Operations) bool {
	valid := true
	AXorB, err := getOpByInputs(ops, a.A, a.B, "XOR")
	if err != nil {
		fmt.Printf("%s XOR %s not found\n", a.A, a.B)
		valid = false
	}
	AAndB, err := getOpByInputs(ops, a.A, a.B, "AND")
	if err != nil {
		fmt.Printf("%s AND %s not found\n", a.A, a.B)
		valid = false
	}
	if len(a.CarryIn) == 0 {
		fmt.Println("CarryIn not present, inferring from output")
		a.CarryIn, err = getOpByIO(ops, AXorB, a.Sum, "XOR")
		if err != nil {
			fmt.Printf("%s XOR ? -> %s not found\n", AXorB, a.Sum)
		}
	} else {
		AXorBXorC, err := getOpByInputs(ops, AXorB, a.CarryIn, "XOR")
		if err != nil {
			fmt.Printf("%s XOR %s not found\n", AXorB, a.CarryIn)
			valid = false
		}
		if AXorBXorC != a.Sum {
			fmt.Printf("%s XOR %s got %s, want %s\n", AXorB, a.CarryIn, AXorBXorC, a.Sum)
			valid = false
		}
	}
	AXorBAndC, err := getOpByInputs(ops, AXorB, a.CarryIn, "AND")
	if err != nil {
		fmt.Printf("%s AND %s not found\n", AXorB, a.CarryIn)
		valid = false
	}
	a.CarryOut, err = getOpByInputs(ops, AXorBAndC, AAndB, "OR")
	if err != nil {
		fmt.Printf("%s OR %s not found\n", AXorBAndC, AAndB)
		valid = false
	}
	// if AXorBAndCOrAAndB != a.CarryOut {
	// 	fmt.Printf("%s OR %s got %s, want %s\n", AXorBAndC, AAndB, AXorBAndCOrAAndB, a.CarryOut)
	// 	valid = false
	// }
	return valid
}

func getOpByInputs(ops Operations, A, B, Op string) (string, error) {
	if len(A) == 0 || len(B) == 0 {
		return "", errors.New("invalid input")
	}
	for out, op := range ops {
		if op.op != Op {
			continue
		}
		if (op.a == A && op.b == B) || (op.a == B && op.b == A) {
			return out, nil
		}
	}
	return "", errors.New("operation not found")
}

func getOpByIO(ops Operations, A, Out, Op string) (string, error) {
	if len(A) == 0 || len(Op) == 0 {
		return "", errors.New("invalid input")
	}
	op, ok := ops[Out]
	if !ok {
		return "", errors.New("operation not found")
	}
	if op.op != Op {
		return "", errors.New("operation not found")
	}
	if op.a != A && op.b != A || op.a == op.b {
		return "", errors.New("operation not found")
	}
	if op.a == A {
		return op.b, nil
	}
	return op.a, nil
}
