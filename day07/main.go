package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type folder struct {
	name     string
	contents []fileItem
	parent   *folder
	children []*folder
}

type fileItem struct {
	name string
	size int
}

func part1(root *folder, limit int) (ans int) {
	sizes := make([]int, 0)
	dfs(root, &sizes)

	for _, size := range sizes {
		if size <= limit {
			ans += size
		}
	}
	return ans
}

func part2(root *folder, space int, required int) (ans int) {
	diff := math.MaxInt
	sizes := make([]int, 0)
	dfs(root, &sizes)

	used := sizes[len(sizes)-1]
	unused := space - used
	goal := required - unused

	for _, size := range sizes {
		if (size >= goal) && (size-goal < diff) {
			diff = size - goal
			ans = size
		}
	}
	return ans
}

func dfs(root *folder, sizes *[]int) int {
	curr := 0

	for _, f := range root.contents {
		curr += f.size
	}

	for _, child := range root.children {
		curr += dfs(child, sizes)
	}

	*sizes = append(*sizes, curr)
	return curr
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var root folder
	root.name = "root"
	curr := &root

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$ cd") {
			var dir string
			fmt.Sscanf(line, "$ cd %s", &dir)

			if dir == ".." {
				curr = curr.parent
			} else if dir != "/" {
				for _, child := range curr.children {
					if child.name == dir {
						curr = child
						break
					}
				}
			}
		} else if strings.HasPrefix(line, "dir") {
			var dir string
			var child folder

			fmt.Sscanf(line, "dir %s", &dir)
			child.name = dir
			child.parent = curr
			curr.children = append(curr.children, &child)
		} else if line != "$ ls" {
			var size int
			var name string

			fmt.Sscanf(line, "%d %s", &size, &name)
			f := fileItem{name, size}
			curr.contents = append(curr.contents, f)
		}
	}
	fmt.Println("Part1:", part1(&root, 100000))
	fmt.Println("Part2:", part2(&root, 70000000, 30000000))
}
