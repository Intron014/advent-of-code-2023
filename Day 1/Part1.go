package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
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
