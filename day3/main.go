package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const filename = "input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}

func part1(file string) int {
	sum := 0
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	for y, l := range lines {
		pn := ""
		valid := false
		l = strings.TrimSpace(l) + "."

		for x := 0; x < len(l); x++ {
			if isNum(l[x]) {
				pn += string(l[x])
				if checkAdjacent(lines, x, y, isSym) {
					valid = true
				}
			} else {
				if valid && pn != "" {
					n, _ := strconv.Atoi(pn)
					sum += n
				}
				pn = ""
				valid = false
			}
		}
	}
	return sum
}

func part2(file string) int {
	sum := 0
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")

	for y, l := range lines {
		l = strings.TrimSpace(l) + "."

		for x := 0; x < len(l); x++ {
			if byte('*') == l[x] {
				n := gearRatio(lines, x, y)
				sum += n
			}
		}
	}
	return sum
}

func isNum(c byte) bool {
	return (c - 0x30) <= 9
}

func isSym(c byte) bool {
	if !isNum(c) && c != 46 {
		return true
	}
	return false
}

// build a square around the coord
var square = [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

func checkAdjacent(l []string, x, y int, checkFunc func(byte) bool) bool {
	for _, d := range square {
		dx, dy := d[0], d[1]
		if y+dy > 0 && y+dy < len(l) && x+dx > 0 && x+dx < len(l[y+dy]) {
			if checkFunc(l[y+dy][x+dx]) {
				return true
			}
		}
	}

	return false
}

func numberValue(s string, x int) int {
	if x < 0 || x >= len(s) || !isNum(s[x]) {
		return -1
	}

	pn := string(s[x])
	for i := 1; x+i < len(s) && isNum(s[x+i]); i++ {
		pn += string(s[x+i])
	}
	for i := 1; x-i >= 0 && isNum(s[x-i]); i++ {
		pn = string(s[x-i]) + pn
	}

	n, _ := strconv.Atoi(pn)
	return n
}

func gearRatio(l []string, x, y int) int {
	n, n1 := 0, 0
	for _, d := range square {
		dx, dy := d[0], d[1]
		if y+dy >= 0 && y+dy < len(l) && x+dx >= 0 && x+dx < len(l[y+dy]) {
			n = numberValue(l[y+dy], x+dx)
			if n > 0 {
				if n1 == 0 {
					n1 = n
					continue
				}
				if n1 != n {
					return n * n1
				}
			}
		}
	}
	return 0
}
