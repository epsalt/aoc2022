package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	add int = iota
	sub
	mul
	div
)

type job struct {
	a, b string
	op   int
}

func collapse(name string, done map[string]int, todo map[string]job) (int, bool) {
	if n, ok := done[name]; ok {
		return n, true
	}
	j := todo[name]

	if j.a == "humn" || j.b == "humn" {
		return -1, false
	}

	a, aok := collapse(j.a, done, todo)
	b, bok := collapse(j.b, done, todo)

	if !aok || !bok {
		return -1, false
	}

	switch j.op {
	case add:
		return a + b, true
	case sub:
		return a - b, true
	case mul:
		return a * b, true
	case div:
		return a / b, true
	}
	panic("bad opcode")
}

func solve(name string, goal int, done map[string]int, todo map[string]job) int {
	if name == "humn" {
		return goal
	}

	j := todo[name]
	a, _ := collapse(j.a, done, todo)
	b, bval := collapse(j.b, done, todo)

	var newGoal int
	if bval {
		switch j.op {
		case add:
			newGoal = goal - b
		case sub:
			newGoal = goal + b
		case mul:
			newGoal = goal / b
		case div:
			newGoal = goal * b
		}
		return solve(j.a, newGoal, done, todo)
	}

	switch j.op {
	case add:
		newGoal = goal - a
	case sub:
		newGoal = a - goal
	case mul:
		newGoal = goal / a
	case div:
		newGoal = b / goal
	}
	return solve(j.b, newGoal, done, todo)
}

func part1(name string, done map[string]int, todo map[string]job) (ans int) {
	if n, ok := done[name]; ok {
		return n
	}

	j := todo[name]
	a := part1(j.a, done, todo)
	b := part1(j.b, done, todo)

	switch j.op {
	case add:
		return a + b
	case sub:
		return a - b
	case mul:
		return a * b
	case div:
		return a / b
	}
	panic("bad opcode")
}

func part2(done map[string]int, todo map[string]job) (ans int) {
	root := todo["root"]
	v := part1(root.b, done, todo)
	return solve(root.a, v, done, todo)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	done := make(map[string]int)
	todo := make(map[string]job)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var name string
		var val int

		_, err := fmt.Sscanf(line, "%s %d", &name, &val)
		if err == nil {
			done[name[:len(name)-1]] = val
			continue
		}

		var a, b, op string
		var opcode int

		fmt.Sscanf(line, "%s %s %s %s", &name, &a, &op, &b)

		switch op {
		case "+":
			opcode = 0
		case "-":
			opcode = 1
		case "*":
			opcode = 2
		case "/":
			opcode = 3
		}

		todo[name[:len(name)-1]] = job{a, b, opcode}
	}
	fmt.Println("Part1:", part1("root", done, todo))
	fmt.Println("Part2:", part2(done, todo))
}
