package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

var blueRe = regexp.MustCompile(`(\d+) blue`)
var redRe = regexp.MustCompile(`(\d+) red`)
var greenRe = regexp.MustCompile(`(\d+) green`)

func main() {
	readFile, _ := os.Open("input.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		input := fileScanner.Text()
		colors := validateLine(input)
		sum += (colors[0] * colors[1] * colors[2])
	}

	log.Println(sum)
}

func validateLine(input string) []int {
	minBlue := 0
	minRed := 0
	minGreen := 0

	blueMatches := blueRe.FindAllStringSubmatch(input, -1)
	for _, match := range blueMatches {
		val, _ := strconv.Atoi(match[1])
		if val > minBlue {
			minBlue = val
		}
	}

	redMatches := redRe.FindAllStringSubmatch(input, -1)
	for _, match := range redMatches {
		val, _ := strconv.Atoi(match[1])
		if val > minRed {
			minRed = val
		}
	}

	greenMatches := greenRe.FindAllStringSubmatch(input, -1)
	for _, match := range greenMatches {
		val, _ := strconv.Atoi(match[1])
		if val > minGreen {
			minGreen = val
		}
	}

	return []int{minRed, minBlue, minGreen}
}
