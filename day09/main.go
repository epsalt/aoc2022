package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type move struct {
	direction string
	distance  int
}

type loc struct {
	x, y int
}

func (tail *loc) advance(head *loc) {
	if (math.Abs(float64(head.x-tail.x)) <= 1) &&
		(math.Abs(float64(head.y-tail.y)) <= 1) {
		return
	}
	dx, dy := 1, 1
	if head.y < tail.y {
		dy = -1
	}
	if head.x < tail.x {
		dx = -1
	}
	if head.x == tail.x {
		tail.y += dy
	} else if head.y == tail.y {
		tail.x += dx
	} else {
		tail.x += dx
		tail.y += dy
	}
}

func part1(moves []move) (ans int) {
	head, tail := loc{0, 0}, loc{0, 0}
	m := make(map[loc]bool)
	m[tail] = true

	for _, move := range moves {
		for i := 0; i < move.distance; i++ {
			switch move.direction {
			case "U":
				head.y += 1
			case "D":
				head.y -= 1
			case "R":
				head.x += 1
			case "L":
				head.x -= 1
			}
			tail.advance(&head)
			m[tail] = true
		}
	}
	return len(m)
}

func part2(moves []move) (ans int) {
	knots := [10]loc{}
	m := make(map[loc]bool)
	m[knots[9]] = true

	for _, move := range moves {
		for i := 0; i < move.distance; i++ {
			switch move.direction {
			case "U":
				knots[0].y += 1
			case "D":
				knots[0].y -= 1
			case "R":
				knots[0].x += 1
			case "L":
				knots[0].x -= 1
			}
			for i := 1; i < len(knots); i++ {
				knots[i].advance(&knots[i-1])
			}
			m[knots[9]] = true
		}
	}
	return len(m)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	moves := make([]move, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var direction string
		var distance int
		line := scanner.Text()

		fmt.Sscanf(line, "%s %d", &direction, &distance)
		moves = append(moves, move{direction, distance})
	}
	fmt.Println("Part1:", part1(moves))
	fmt.Println("Part2:", part2(moves))
}
