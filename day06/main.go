package main

import (
	"fmt"
	"log"
	"os"
)

func marker(code []byte, n int) int {
	m := make(map[byte]int)

	for i := 0; i < len(code); i++ {
		m[code[i]] += 1

		if i >= n {
			m[code[i-n]] -= 1

			if m[code[i-n]] == 0 {
				delete(m, code[i-n])
			}
		}

		if len(m) == n {
			return i + 1
		}
	}
	return len(code) + 1
}

func main() {
	code, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part1:", marker(code, 4))
	fmt.Println("Part2:", marker(code, 14))
}
