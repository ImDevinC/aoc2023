package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)

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
	readFile, _ := os.Open("input1.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	calibration := 0
	for fileScanner.Scan() {
		input := fileScanner.Text()
		val := getValues(input)
		fmt.Printf("Input: %s, Value: %d\n", input, val)
		calibration += val
	}
	fmt.Println(calibration)
}

func getValues(input string) int {
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
