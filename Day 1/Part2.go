package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, _ := os.Open("calibration.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	regex := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|[0-9]`)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)

		matches := regex.FindAllString(line, -1)
		fmt.Println("Matches", matches)
		if len(matches) > 0 {
			first := matches[0]
			firstDigit := 0
			if val, ok := numberMap[first]; ok {
				firstDigit = val
			} else {
				firstDigit = int(first[0] - '0')
			}

			last := matches[len(matches)-1]
			lastDigit := 0
			if val, ok := numberMap[last]; ok {
				lastDigit = val
			} else {
				lastDigit = int(last[0] - '0')
			}
			fmt.Print("F ", firstDigit)
			fmt.Print(" L ", lastDigit)
			sum += firstDigit*10 + lastDigit
			fmt.Println(" T", sum)
		}
	}

	fmt.Println(sum)
}
