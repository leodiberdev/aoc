package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TwoDigits struct {
	first      int
	firstIndex int
	last       int
	lastIndex  int
}

type SpelledNumber struct {
	first      string
	firstIndex int
	last       string
	lastIndex  int
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func isSpelledDigit(s string) bool {
	spelledDigits := "one-two-three-four-five-six-seven-eight-nine"
	return strings.Contains(spelledDigits, s)
}

func returnCalibration(line string) (TwoDigits, SpelledNumber) {
	var number TwoDigits
	var spelledNumber SpelledNumber
	// var currentSubstring string

	for i, c := range line {
		if isDigit(c) {
			if number.first == 0 {
				number.first = int(c-'0') * 10
				number.firstIndex = i
			}
			number.last = int(c - '0')
			number.lastIndex = i
		}

	}
	return number, spelledNumber
}

func main() {
	input_path := "../inputs/day01/pt01-example"
	var sum int

	file, err := os.Open(input_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		number, _ := returnCalibration(line)

		sum += number.first + number.last
	}

	fmt.Println(sum)
}
