package main

import (
	"fmt"
	"strconv"
	"strings"
)

var numbersMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seve":  7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

func findNumberInLine(line string) []int {
	var err error
	var numbas []int
	text := ""
	number := 0
	ok := false
	for _, _rune := range line {
		number, err = strconv.Atoi(string(_rune))
		if err != nil {
			text += string(_rune)
			for key := range numbersMap {
				if strings.Contains(text, key) {
					text = key
				}
			}
			number, ok = numbersMap[text]
		} else {
			ok = true
		}
		if ok {
			numbas = append(numbas, number)
			if len(text) > 1 {
				text = string(text[len(text)-1])
			}
		}
	}

	return numbas
}

func run2() {
	content := readFile()
	total := 0
	for _, line := range strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n") {
		digitArray := findNumberInLine(line)
		number := twoDigitNumber(digitArray)
		total += number
	}
	fmt.Println(total)
}
