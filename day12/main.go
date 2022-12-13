package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type loc struct{ x, y int }

func part1(grid [][]int, start, end loc) (ans int) {
	r, c := len(grid), len(grid[0])
	dists := make([][]int, 0)

	for i := 0; i < r; i++ {
		row := make([]int, c)
		for j := 0; j < c; j++ {
			row[j] = math.MaxInt
		}
		dists = append(dists, row)
	}
	dists[start.x][start.y] = 0

	q := []loc{start}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, n := range []loc{{curr.x + 1, curr.y}, {curr.x - 1, curr.y}, {curr.x, curr.y + 1}, {curr.x, curr.y - 1}} {
			if (n.x >= 0 && n.x < r) && (n.y >= 0 && n.y < c) && (grid[n.x][n.y] <= grid[curr.x][curr.y]+1) {
				if dists[n.x][n.y] > dists[curr.x][curr.y]+1 {
					q = append(q, n)
					dists[n.x][n.y] = dists[curr.x][curr.y] + 1
				}
			}
		}
	}
	return dists[end.x][end.y]
}

func part2(grid [][]int, end loc) (ans int) {
	r, c := len(grid), len(grid[0])
	starts := make([]loc, 0)
	ans = math.MaxInt

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 0 {
				starts = append(starts, loc{i, j})
			}
		}
	}
	for _, start := range starts {
		steps := part1(grid, start, end)
		if steps < ans {
			ans = steps
		}
	}
	return ans
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make([][]int, 0)
	var start, end loc
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)
		for j, r := range line {
			if r == 'S' {
				row = append(row, 0)
				start = loc{i, j}
			} else if r == 'E' {
				row = append(row, 25)
				end = loc{i, j}
			} else {
				row = append(row, int(r-'a'))
			}
		}
		grid = append(grid, row)
		i++

	}
	fmt.Println("Part1:", part1(grid, start, end))
	fmt.Println("Part2:", part2(grid, end))
}
