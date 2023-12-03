package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	gameIDSum := 0
	powerSum := 0
	fileLocation := os.Getenv("DAY2_FILE")
	if fileLocation == "" {
		fallback := "day-2/day2.txt"
		fmt.Printf("Falling back to %v as james is a dummy and didnt give me a file \n", fallback)
		fileLocation = fallback
	}

	file, err := os.Open(fileLocation)

	if err != nil {
		fmt.Printf("Error finding file %v", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := lineSplitter(scanner.Text())

		fmt.Println(line)
		splitLines, power := ballCounter(line[1:])
		if splitLines {
			gameNumber := strings.TrimPrefix(line[0], "Game")
			gameNumberAsInt, _ := strconv.Atoi(strings.TrimSpace(gameNumber))
			gameIDSum = gameIDSum + gameNumberAsInt
		}
		powerSum = power + powerSum
	}

	fmt.Printf("GameID sum: %v \n", gameIDSum)
	fmt.Printf("PowerSum: %v \n", powerSum)
}

// index 0, game name, Rest = ball sets
func lineSplitter(line string) []string {

	//Get the ID
	idSplit := strings.Split(line, ":")
	rounds := strings.Split(idSplit[1], ";")

	// +1 to add the ID.
	parsedLine := make([]string, 1)
	parsedLine[0] = idSplit[0]

	return append(parsedLine, rounds...)

}

// Error if a single pull shows > 13 balls
func ballCounter(game []string) (validGame bool, power int) {
	var red, green, blue int

	highestRed := 1
	highestBlue := 1
	highestGreen := 1

	valid := true

	for _, round := range game {

		for _, item := range strings.Split(round, ",") {
			switch {
			case strings.HasSuffix(item, "red"):
				x := strings.TrimSuffix(item, "red")
				conversion, _ := strconv.Atoi(strings.TrimSpace(x))
				red = red + conversion
			case strings.HasSuffix(item, "blue"):
				x := strings.TrimSuffix(item, "blue")
				conversion, _ := strconv.Atoi(strings.TrimSpace(x))
				blue = blue + conversion
			case strings.HasSuffix(item, "green"):
				x := strings.TrimSuffix(item, "green")
				conversion, _ := strconv.Atoi(strings.TrimSpace(x))
				green = green + conversion
			}

		}

		if red > 12 || green > 13 || blue > 14 {
			valid = false
		}

		if highestBlue < blue {
			highestBlue = blue
		}

		if highestRed < red {
			highestRed = red
		}

		if highestGreen < green {
			highestGreen = green
		}

		red = 0
		green = 0
		blue = 0
	}
	return valid, highestBlue * highestRed * highestGreen
}
