package main

import (
	"bufio"
	"fmt"
	"log"
	// "math"
	"os"
	"strings"
)

type valve struct {
	name    string
	rate    int
	tunnels []string
}

type state struct {
	pos  int
	open bitset
	time int
}

type state2 struct {
	pos1  int
	pos2  int
	open  bitset
	time1 int
	time2 int
}

type edge struct {
	u, v int
}

type bitset uint

func (b bitset) set(pos uint) bitset {
	b |= (1 << pos)
	return b
}

func (b bitset) has(pos uint) bool {
	val := b & (1 << pos)
	return (val > 0)
}

func part1(start int, valves []valve, m map[string]int, weights map[edge]int, time int) (ans int) {
	var open bitset

	var helper func(pos int, open bitset, time int) (pressure int)
	cache := make(map[state]int)

	helper = func(pos int, open bitset, time int) (pressure int) {
		if a, ok := cache[state{pos, open, time}]; ok {
			return a
		}

		pressures := make([]int, 0)

		for i := 0; i < len(valves); i++ {
			if valves[i].rate == 0 {
				continue
			}

			dist := weights[edge{pos, i}]

			if !open.has(uint(i)) && dist < time {
				remaining := valves[i].rate * (time - dist - 1)
				newOpen := open.set(uint(i))
				pressures = append(pressures, helper(i, newOpen, time-dist-1)+remaining)
			}
		}

		maxPressure := 0
		for _, p := range pressures {
			if p > maxPressure {
				maxPressure = p
			}
		}

		cache[state{pos, open, time}] = maxPressure
		return maxPressure
	}
	return helper(start, open, time)
}

func part2(start int, valves []valve, m map[string]int, weights map[edge]int, time int) (ans int) {
	var open bitset

	var helper func(pos1 int, pos2 int, open bitset, time1 int, time2 int) (pressure int)
	cache := make(map[state2]int)

	helper = func(pos1 int, pos2 int, open bitset, time1 int, time2 int) (pressure int) {
		if a, ok := cache[state2{pos1, pos2, open, time1, time2}]; ok {
			return a
		}

		pressures := make([]int, 0)

		for i := 0; i < len(valves); i++ {
			if valves[i].rate == 0 {
				continue
			}

			if time1 > time2 {
				dist := weights[edge{pos1, i}]
				if !open.has(uint(i)) && dist < time1 {
					remaining := valves[i].rate * (time1 - dist - 1)
					newOpen := open.set(uint(i))
					pressures = append(pressures, helper(i, pos2, newOpen, time1-dist-1, time2)+remaining)
				}
			} else {
				dist := weights[edge{pos2, i}]
				if !open.has(uint(i)) && dist < time2 {
					remaining := valves[i].rate * (time2 - dist - 1)
					newOpen := open.set(uint(i))
					pressures = append(pressures, helper(pos1, i, newOpen, time1, time2-dist-1)+remaining)
				}
			}
		}

		maxPressure := 0
		for _, p := range pressures {
			if p > maxPressure {
				maxPressure = p
			}
		}

		cache[state2{pos1, pos2, open, time1, time2}] = maxPressure
		return maxPressure
	}
	return helper(start, start, open, time, time)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	valves := make([]valve, 0)
	m := make(map[string]int)
	i := 0
	var start int

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "; ")

		var name string
		var rate int

		fmt.Sscanf(arr[0], "Valve %s has flow rate=%d", &name, &rate)
		if name == "AA" {
			start = i
		}

		tunnels := strings.Split(strings.TrimSpace(arr[1][22:]), ", ")
		v := valve{name, rate, tunnels}
		m[name] = i
		valves = append(valves, v)
		i++
	}

	weights := make(map[edge]int)

	for i := range valves {
		for j := range valves {
			if i == j {
				weights[edge{i, i}] = 0
			} else {
				weights[edge{i, j}] = 9999
			}
			for _, tunnel := range valves[i].tunnels {
				if m[tunnel] == j {
					weights[edge{i, j}] = 1
				}
			}
		}
	}

	for k := range valves {
		for i := range valves {
			for j := range valves {
				if weights[edge{i, j}] > (weights[edge{i, k}] + weights[edge{k, j}]) {
					weights[edge{i, j}] = (weights[edge{i, k}] + weights[edge{k, j}])
				}
			}
		}
	}

	for k, v := range weights {
		weights[edge{k.v, k.u}] = v
	}

	for k, v := range weights {
		if (k.u != start && valves[k.u].rate == 0) || valves[k.v].rate == 0 || v == 0 {
			delete(weights, k)
		}
	}

	fmt.Println("Part1:", part1(start, valves, m, weights, 30))
	fmt.Println("Part2:", part2(start, valves, m, weights, 26))
}
