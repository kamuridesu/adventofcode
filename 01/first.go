package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile() string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func findNumberInLineOld(line string) []int {
	var numbers []int
	for _, _rune := range line {
		number, err := strconv.Atoi(string(_rune))
		if err == nil {
			numbers = append(numbers, number)
		}
	}
	return numbers
}

func twoDigitNumber(digits []int) int {
	first := digits[0]
	var last int
	if len(digits) < 2 {
		last = digits[0]
	} else {
		last = digits[len(digits)-1]
	}
	n, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func run() {
	content := readFile()
	total := 0
	for _, line := range strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n") {
		digitArray := findNumberInLineOld(line)
		number := twoDigitNumber(digitArray)
		total += number
	}
	fmt.Println(total)
}
