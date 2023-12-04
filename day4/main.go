package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const filename = "input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}

func part1(file string) int {
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	sum := 0
	for _, l := range lines {
		val, _ := checkCard(l)
		sum += val
	}
	return sum
}

func part2(file string) int {
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	cards := map[int]int{}
	sum := 0
	for i, l := range lines {
		id := i + 1
		extras := checkExtraCards(l)
		cards[id] += 1
		multiplier := 1
		if val, exists := cards[id]; exists {
			multiplier = val
		}
		for j := 1; j <= extras; j++ {
			cards[j+id] += 1 * multiplier
		}
	}
	for _, v := range cards {
		sum += v
	}
	return sum
}

func checkExtraCards(l string) int {
	_, wins := checkCard(l)
	return wins
}

func checkCard(l string) (int, int) {
	vals := strings.Split(l, ":")[1]
	numbers := strings.Split(strings.TrimSpace(strings.Split(vals, "|")[0]), " ")
	results := strings.Split(strings.TrimSpace(strings.Split(vals, "|")[1]), " ")
	score := 0
	winsFound := map[string]bool{}
	wins := 0
	for _, r := range results {
		for _, n := range numbers {
			res := strings.TrimSpace(r)
			num := strings.TrimSpace(n)
			if len(num) > 0 && num == res && !winsFound[num] {
				winsFound[num] = true
				wins++
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
	}
	return score, wins
}
