package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var rePart1 = regexp.MustCompile(`\d`)
var rePart2 = regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)

var words = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: go run main.go <part1|part2>")
		return
	}

	re := &regexp.Regexp{}

	if os.Args[1] == "part1" {
		re = rePart1
	} else if os.Args[1] == "part2" {
		re = rePart2
	} else {
		log.Println("Usage: go run main.go <part1|part2>")
		return
	}

	readFile, _ := os.Open("input1.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	calibration := 0
	for fileScanner.Scan() {
		input := fileScanner.Text()
		val := getValues(re, input)
		fmt.Printf("Input: %s, Value: %d\n", input, val)
		calibration += val
	}
	fmt.Println(calibration)
}

func getValues(re *regexp.Regexp, input string) int {
	match := re.FindString(input)
	var value string
	if val, exists := words[match]; exists {
		value = val
	} else {
		value = match
	}

	var final string
	for i := 0; i < len(input); i++ {
		match := re.FindString(input[i:])
		if match != "" {
			final = match
		}
	}

	if val, exists := words[final]; exists {
		value += val
	} else {
		value += final
	}

	combined, _ := strconv.Atoi(value)

	return combined
}
