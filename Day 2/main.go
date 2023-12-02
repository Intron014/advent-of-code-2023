package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	ID     int
	Rounds []map[string]int
}

func main() {
	part1()
	part2()
}

func part1() {
	games := parseInput()

	sum := 0
	for _, game := range games {
		if isPossible(game) {
			sum += game.ID
		}
	}

	fmt.Println(sum)
}

func part2() {
	games := parseInput()

	sum := 0
	for _, game := range games {
		power := minCubesPower(game)
		sum += power
	}

	fmt.Println(sum)
}

func parseInput() []Game {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	gameRegex := regexp.MustCompile(`Game (\d+): (.*)`)
	cubeRegex := regexp.MustCompile(`(\d+) (\w+)`)

	var games []Game
	for scanner.Scan() {
		line := scanner.Text()
		gameMatch := gameRegex.FindStringSubmatch(line)
		gameID, _ := strconv.Atoi(gameMatch[1])
		roundsStr := strings.Split(gameMatch[2], "; ")

		var rounds []map[string]int
		for _, roundStr := range roundsStr {
			round := make(map[string]int)
			cubeMatches := cubeRegex.FindAllStringSubmatch(roundStr, -1)
			for _, cubeMatch := range cubeMatches {
				count, _ := strconv.Atoi(cubeMatch[1])
				color := cubeMatch[2]
				round[color] += count
			}
			rounds = append(rounds, round)
		}

		games = append(games, Game{ID: gameID, Rounds: rounds})
	}

	return games
}

func isPossible(game Game) bool {
	for _, round := range game.Rounds {
		if round["red"] > 12 || round["green"] > 13 || round["blue"] > 14 {
			return false
		}
	}
	return true
}

func minCubesPower(game Game) int {
	minCubes := map[string]int{"red": 0, "green": 0, "blue": 0}

	for _, round := range game.Rounds {
		for color, count := range round {
			if count > minCubes[color] {
				minCubes[color] = count
			}
		}
	}

	return minCubes["red"] * minCubes["green"] * minCubes["blue"]
}
