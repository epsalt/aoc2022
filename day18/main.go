package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type cube struct {
	x, y, z int
}

func (c cube) String() string {
	return fmt.Sprintf("%d,%d,%d", c.x, c.y, c.z)
}

type limit struct {
	min, max int
}

type cubeLimit struct {
	x, y, z limit
}

func part1(cubes []cube) (ans int) {
	m := make(map[cube]bool)

	for _, c := range cubes {
		m[c] = true
	}

	for _, c := range cubes {
		for _, d := range []int{-1, 1} {
			if !m[cube{c.x + d, c.y, c.z}] {
				ans += 1
			}
			if !m[cube{c.x, c.y + d, c.z}] {
				ans += 1
			}
			if !m[cube{c.x, c.y, c.z + d}] {
				ans += 1
			}
		}
	}
	return ans
}

func part2(cubes []cube) (ans int) {
	m := make(map[cube]bool)

	for _, c := range cubes {
		m[c] = true
	}

	empty := make([]cube, 0)
	empties := make(map[string]int)
	limits := cubeLimit{limit{math.MaxInt, 0}, limit{math.MaxInt, 0}, limit{math.MaxInt, 0}}

	for _, c := range cubes {
		if c.x < limits.x.min {
			limits.x.min = c.x
		}
		if c.x > limits.x.max {
			limits.x.max = c.x
		}

		if c.y < limits.y.min {
			limits.y.min = c.y
		}
		if c.y > limits.y.max {
			limits.y.max = c.y
		}

		if c.z < limits.z.min {
			limits.z.min = c.z
		}
		if c.z > limits.z.max {
			limits.z.max = c.z
		}

		for _, d := range []int{-1, 1} {
			if !m[cube{c.x + d, c.y, c.z}] {
				empty = append(empty, cube{c.x + d, c.y, c.z})
				empties[cube{c.x + d, c.y, c.z}.String()] += 1
			}
			if !m[cube{c.x, c.y + d, c.z}] {
				empty = append(empty, cube{c.x, c.y + d, c.z})
				empties[cube{c.x, c.y + d, c.z}.String()] += 1
			}
			if !m[cube{c.x, c.y, c.z + d}] {
				empty = append(empty, cube{c.x, c.y, c.z + d})
				empties[cube{c.x, c.y, c.z + d}.String()] += 1
			}
		}
	}

	for _, e := range empty {
		visited := make(map[cube]bool)
		if !internal(e, m, visited, limits) {
			ans += 1
		}
	}
	return ans
}

func internal(c cube, m map[cube]bool, visited map[cube]bool, limits cubeLimit) (ok bool) {
	ok = true
	if visited[c] {
		return true
	}
	visited[c] = true

	if c.x < limits.x.min || c.x > limits.x.max || c.y < limits.y.min || c.y > limits.y.max || c.z < limits.z.min || c.z > limits.z.max {
		return false
	}
	for _, d := range []int{-1, 1} {
		if !m[cube{c.x + d, c.y, c.z}] {
			ok = ok && internal(cube{c.x + d, c.y, c.z}, m, visited, limits)
		}

		if !m[cube{c.x, c.y + d, c.z}] {
			ok = ok && internal(cube{c.x, c.y + d, c.z}, m, visited, limits)
		}
		if !m[cube{c.x, c.y, c.z + d}] {
			ok = ok && internal(cube{c.x, c.y, c.z + d}, m, visited, limits)
		}
	}
	return ok
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cubes := make([]cube, 0)

	for scanner.Scan() {
		line := scanner.Text()
		var x, y, z int

		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)

		cubes = append(cubes, cube{x, y, z})
	}

	fmt.Println("Part1:", part1(cubes))
	fmt.Println("Part2:", part2(cubes))
}
