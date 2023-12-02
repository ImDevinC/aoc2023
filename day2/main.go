package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

var valid = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var blueRe = regexp.MustCompile(`(\d+) blue`)
var redRe = regexp.MustCompile(`(\d+) red`)
var greenRe = regexp.MustCompile(`(\d+) green`)

func main() {
	readFile, _ := os.Open("input.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	validCount := 0
	id := 0
	for fileScanner.Scan() {
		id++
		input := fileScanner.Text()
		if validateLine(input) {
			validCount += id
		}
	}
	log.Println("Valid pulls:", validCount)
}

func validateLine(input string) bool {
	blueMatches := blueRe.FindAllStringSubmatch(input, -1)
	for _, match := range blueMatches {
		val, _ := strconv.Atoi(match[1])
		if val > valid["blue"] {
			return false
		}
	}

	redMatches := redRe.FindAllStringSubmatch(input, -1)
	for _, match := range redMatches {
		val, _ := strconv.Atoi(match[1])
		if val > valid["red"] {
			return false
		}
	}

	greenMatches := greenRe.FindAllStringSubmatch(input, -1)
	for _, match := range greenMatches {
		val, _ := strconv.Atoi(match[1])
		if val > valid["green"] {
			return false
		}
	}

	return true
}
