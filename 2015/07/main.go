package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Expr struct {
	Op   string   // "", "AND", "OR", "LSHIFT", "RSHIFT", "NOT"
	Args []string // wires or numbers
}

func main() {
	file, err := helpers.GetFile("2015/07/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	input, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	wireExpr := make(map[string]Expr)

	for _, line := range lines {
		fmt.Println(lines)
		parts := strings.Split(line, " -> ")
		left := parts[0]
		target := parts[1]

		tokens := strings.Split(left, " ")

		// Parse expression
		switch len(tokens) {
		case 1:
			// "123" -> x   OR   "ab" -> x
			wireExpr[target] = Expr{Op: "SET", Args: []string{tokens[0]}}
		case 2:
			// "NOT x" -> y
			wireExpr[target] = Expr{Op: "NOT", Args: []string{tokens[1]}}
		case 3:
			// "x AND y" -> z
			// "p LSHIFT 2" -> q
			wireExpr[target] = Expr{Op: tokens[1], Args: []string{tokens[0], tokens[2]}}
		}
	}

	cache := make(map[string]uint16)

	var eval func(string) uint16
	eval = func(sym string) uint16 {
		// Is constant number?
		if v, err := strconv.Atoi(sym); err == nil {
			return uint16(v)
		}

		// Already computed?
		if v, ok := cache[sym]; ok {
			return v
		}

		expr := wireExpr[sym]
		var result uint16

		switch expr.Op {
		case "SET":
			result = eval(expr.Args[0])
		case "NOT":
			result = ^eval(expr.Args[0])
		case "AND":
			result = eval(expr.Args[0]) & eval(expr.Args[1])
		case "OR":
			result = eval(expr.Args[0]) | eval(expr.Args[1])
		case "LSHIFT":
			n, _ := strconv.Atoi(expr.Args[1])
			result = eval(expr.Args[0]) << n
		case "RSHIFT":
			n, _ := strconv.Atoi(expr.Args[1])
			result = eval(expr.Args[0]) >> n

		}

		cache[sym] = result
		return result
	}

	fmt.Println(eval("a"))
}
