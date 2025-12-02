package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	op   string   // "", "NOT", "AND", "OR", "LSHIFT", "RSHIFT"
	args []string // 1 or 2 args (could be wire names or numeric literals)
}

func parseInput(path string) map[string]Instruction {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	inst := make(map[string]Instruction)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, "->")
		if len(parts) != 2 {
			panic("invalid line (no ->): " + line)
		}
		expr := strings.TrimSpace(parts[0])
		out := strings.TrimSpace(parts[1])
		tokens := strings.Fields(expr)

		switch len(tokens) {
		case 1:
			inst[out] = Instruction{op: "", args: []string{tokens[0]}}
		case 2:
			// "NOT x"
			inst[out] = Instruction{op: tokens[0], args: []string{tokens[1]}}
		case 3:
			// "x AND y" or "x LSHIFT 2"
			inst[out] = Instruction{op: tokens[1], args: []string{tokens[0], tokens[2]}}
		default:
			panic("unrecognized instruction: " + line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return inst
}

type EvalContext struct {
	inst  map[string]Instruction
	cache map[string]uint16
}

func (ec *EvalContext) evalArg(arg string) (uint16, bool) {
	// numeric literal?
	if v, err := strconv.Atoi(arg); err == nil {
		return uint16(v), true
	}
	return ec.eval(arg)
}

func (ec *EvalContext) eval(wire string) (uint16, bool) {
	// if arg itself is a numeric literal (someone might call eval with literal)
	if v, err := strconv.Atoi(wire); err == nil {
		return uint16(v), true
	}

	// cached?
	if v, ok := ec.cache[wire]; ok {
		return v, true
	}

	inst, exists := ec.inst[wire]
	if !exists {
		// no instruction for this wire
		return 0, false
	}

	var res uint16
	switch inst.op {
	case "":
		v0, ok := ec.evalArg(inst.args[0])
		if !ok {
			return 0, false
		}
		res = v0
	case "NOT":
		v0, ok := ec.evalArg(inst.args[0])
		if !ok {
			return 0, false
		}
		res = ^v0
	case "AND":
		v0, ok0 := ec.evalArg(inst.args[0])
		v1, ok1 := ec.evalArg(inst.args[1])
		if !ok0 || !ok1 {
			return 0, false
		}
		res = v0 & v1
	case "OR":
		v0, ok0 := ec.evalArg(inst.args[0])
		v1, ok1 := ec.evalArg(inst.args[1])
		if !ok0 || !ok1 {
			return 0, false
		}
		res = v0 | v1
	case "LSHIFT":
		v0, ok0 := ec.evalArg(inst.args[0])
		v1, ok1 := ec.evalArg(inst.args[1])
		if !ok0 || !ok1 {
			return 0, false
		}
		// do shift in wider type then truncate to 16-bit to keep semantics predictable
		res = uint16(uint32(v0) << uint(v1))
	case "RSHIFT":
		v0, ok0 := ec.evalArg(inst.args[0])
		v1, ok1 := ec.evalArg(inst.args[1])
		if !ok0 || !ok1 {
			return 0, false
		}
		res = uint16(uint32(v0) >> uint(v1))
	default:
		panic("unknown op: " + inst.op)
	}

	ec.cache[wire] = res
	return res, true
}

// shallow copy map[string]Instruction
func copyInst(m map[string]Instruction) map[string]Instruction {
	c := make(map[string]Instruction, len(m))
	for k, v := range m {
		c[k] = v
	}
	return c
}

func computeSignal(inst map[string]Instruction, target string) (uint16, bool) {
	ec := &EvalContext{inst: inst, cache: make(map[string]uint16)}
	return ec.eval(target)
}

func main() {
	// file and target can be changed as needed
	inst := parseInput("2015/07/input.txt")
	target := "a"

	// PART 1: compute target (e.g., "a")
	a1, ok := computeSignal(inst, target)
	if !ok {
		fmt.Printf("Part 1: could not compute signal for wire %q (wire may not exist or depends on undefined wire)\n", target)
	} else {
		fmt.Printf("Part 1: signal on %q = %d\n", target, a1)

		// PART 2: override 'b' with the signal from part1, reset all other wires, recompute
		inst2 := copyInst(inst)
		// override b to be the literal value of a1
		inst2["b"] = Instruction{op: "", args: []string{strconv.Itoa(int(a1))}}

		a2, ok2 := computeSignal(inst2, target)
		if !ok2 {
			fmt.Printf("Part 2: could not compute signal for wire %q after overriding b\n", target)
		} else {
			fmt.Printf("Part 2: signal on %q after overriding b = %d\n", target, a2)
		}
	}
}
