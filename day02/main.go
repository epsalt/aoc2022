package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type game struct {
	oppo int
	you  int
}

const (
	rock int = 1 + iota
	paper
	scissors
)

const (
	lose int = 1 + iota
	draw
	win
)

func (g game) score() int {
	switch {
	case g.oppo == g.you:
		return g.you + 3
	case (g.oppo == rock && g.you == scissors) ||
		(g.oppo == paper && g.you == rock) ||
		(g.oppo == scissors && g.you == paper):
		return g.you
	default:
		return g.you + 6
	}
}

func (g game) shape() int {
	switch {
	case g.you == draw:
		return g.oppo
	case (g.you == lose && g.oppo == rock) || (g.you == win && g.oppo == paper):
		return scissors
	case (g.you == lose && g.oppo == scissors) || (g.you == win && g.oppo == rock):
		return paper
	default:
		return rock
	}
}

func part1(games []game) (ans int) {
	for _, g := range games {
		ans += g.score()
	}
	return ans
}

func part2(games []game) (ans int) {
	var n game
	var you int

	for _, g := range games {
		you = g.shape()
		n = game{g.oppo, you}
		ans += n.score()
	}
	return ans
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	games := make([]game, 0)
	var line string
	var g game

	youMap := map[byte]int{'X': rock, 'Y': paper, 'Z': scissors}
	oppoMap := map[byte]int{'A': rock, 'B': paper, 'C': scissors}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		g = game{oppoMap[line[0]], youMap[line[2]]}
		games = append(games, g)
	}
	fmt.Println("Part1:", part1(games))
	fmt.Println("Part2:", part2(games))
}
