package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"unicode"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("calibration.txt")
	if err != nil {
		fmt.Println("Error opening:", err)
		return
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, lastDigit := -1, -1
		for _, r := range line {
			if unicode.IsDigit(r) {
				if firstDigit == -1 {
					firstDigit = int(r - '0')
				}
				lastDigit = int(r - '0')
			}
		}
		if firstDigit != -1 && lastDigit != -1 {
			sum += firstDigit*10 + lastDigit
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading:", err)
	}

	fmt.Println("Sum:", sum)
}

func part2() {
	file, _ := os.Open("calibration.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pattern := regexp.MustCompile(`[1-9]|one|two|three|four|five|six|seven|eight|nine`)
	values := map[string]int{
		"1":     1,
		"one":   1,
		"2":     2,
		"two":   2,
		"3":     3,
		"three": 3,
		"4":     4,
		"four":  4,
		"5":     5,
		"five":  5,
		"6":     6,
		"six":   6,
		"7":     7,
		"seven": 7,
		"8":     8,
		"eight": 8,
		"9":     9,
		"nine":  9,
	}

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		firstDigit := values[pattern.FindString(line)]
		var lastDigit int

		for i := len(line) - 1; i >= 0; i-- {
			if digit := pattern.FindString(line[i:]); digit != "" {
				lastDigit = values[digit]
				break
			}
		}

		total += firstDigit*10 + lastDigit
	}

	fmt.Println("Sum:", total)
}
