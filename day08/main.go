package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type loc struct {
	x, y int
}

func score(xy loc, grid [][]int, heights []int) (score int) {
	height := grid[xy.x][xy.y]

	for k := len(heights) - 1; k >= 0; k-- {
		if heights[k] >= height {
			return score + 1
		}
		score += 1
	}
	return score
}

func part1(grid [][]int) (ans int) {
	visible := make(map[loc]bool)
	r, c := len(grid), len(grid[0])

	for i := 0; i < r; i++ {
		tallest := -1
		for j := 0; j < c; j++ {
			if grid[i][j] > tallest {
				visible[loc{i, j}] = true
				tallest = grid[i][j]
			}
		}
		tallest = -1
		for j := c - 1; j >= 0; j-- {
			if grid[i][j] > tallest {
				visible[loc{i, j}] = true
				tallest = grid[i][j]
			}
		}
	}
	for j := 0; j < c; j++ {
		tallest := -1
		for i := 0; i < r; i++ {
			if grid[i][j] > tallest {
				visible[loc{i, j}] = true
				tallest = grid[i][j]
			}
		}
		tallest = -1
		for i := r - 1; i >= 0; i-- {
			if grid[i][j] > tallest {
				visible[loc{i, j}] = true
				tallest = grid[i][j]
			}
		}
	}
	return len(visible)
}

func part2(grid [][]int) (ans int) {
	scores := make(map[loc]int)
	r, c := len(grid), len(grid[0])

	for i := 0; i < r; i++ {
		heights := make([]int, 0)
		for j := 0; j < c; j++ {
			scores[loc{i, j}] = 1
			scores[loc{i, j}] *= score(loc{i, j}, grid, heights)
			heights = append(heights, grid[i][j])
		}
		heights = heights[:0]
		for j := c - 1; j >= 0; j-- {
			scores[loc{i, j}] *= score(loc{i, j}, grid, heights)
			heights = append(heights, grid[i][j])
		}
	}
	for j := 0; j < c; j++ {
		heights := make([]int, 0)
		for i := 0; i < r; i++ {
			scores[loc{i, j}] *= score(loc{i, j}, grid, heights)
			heights = append(heights, grid[i][j])
		}
		heights = heights[:0]
		for i := r - 1; i >= 0; i-- {
			scores[loc{i, j}] *= score(loc{i, j}, grid, heights)
			heights = append(heights, grid[i][j])
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if scores[loc{i, j}] > ans {
				ans = scores[loc{i, j}]
			}
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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))

		for i, c := range line {
			val, _ := strconv.Atoi(string(c))
			row[i] = val
		}
		grid = append(grid, row)
	}
	fmt.Println("Part1:", part1(grid))
	fmt.Println("Part2:", part2(grid))
}
