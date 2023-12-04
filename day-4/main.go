package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type scratchCard struct {
	Number         int
	UserNumbers    []int
	WinningNumbers []int
}

func main() {
	fileLocation := os.Getenv("DAY4_FILE")
	scratchCards := []scratchCard{}
	cardCounter := 0

	if fileLocation == "" {
		fallback := "day4.txt"
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

		str := strings.Split(scanner.Text(), "|")
		winningNumbers := str[1]
		s := strings.Split(str[0], ":")
		UserNumbers := s[1]
		cardNumber, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(s[0], "Card ")))

		if err != nil {
			fmt.Printf("Error getting card number: %v", err)
			os.Exit(1)
		}

		scratch := scratchCard{
			Number:         cardNumber,
			UserNumbers:    stringToIntArray(UserNumbers, " "),
			WinningNumbers: stringToIntArray(winningNumbers, " "),
		}

		scratchCards = append(scratchCards, scratch)

	}

	// part one answer
	score := calculateCardScore(scratchCards)
	fmt.Printf("Part one answer %v \n", score)

	cardCounter = cardCounter + len(scratchCards)
	returnCards := calculateCardReturns(scratchCards, scratchCards)
	cardCounter = cardCounter + len(returnCards)

	for len(returnCards) > 0 {
		returnCards = calculateCardReturns(returnCards, scratchCards)
		cardCounter = cardCounter + len(returnCards)
	}

	// part two answer
	fmt.Println("Part two answer: ", cardCounter)
}

// Part one. Return score
func calculateCardScore(scratchCards []scratchCard) int {
	sumCounter := 0
	for _, card := range scratchCards {
		doubler := false
		sum := 0
		matchingNumbers := 0
		fmt.Println("Checking card: ", card.Number)
		for _, i := range card.UserNumbers {

			for _, j := range card.WinningNumbers {
				if i == j {
					// fmt.Println("Match on card: ", card.Number)
					matchingNumbers++
					if doubler {
						sum = sum * 2
						continue
					}
					sum++
					doubler = true

				}
			}

		}

		sumCounter = sumCounter + sum
	}
	return sumCounter

}

// Part Two, generate a list of cards to return
func calculateCardReturns(scratchCards, oringalSet []scratchCard) []scratchCard {
	returnCards := []scratchCard{}

	for _, card := range scratchCards {
		counter := 0
		for _, i := range card.UserNumbers {
			for _, j := range card.WinningNumbers {
				if i == j {
					counter++
				}
			}
		}

		for i := 0; i < counter; i++ {
			// fmt.Println("adding card: ", oringalSet[card.Number+i], "from card:", card.Number)
			returnCards = append(returnCards, oringalSet[card.Number+i])
		}
	}
	return returnCards
}

func stringToIntArray(s string, sep string) []int {
	numbers := []int{}

	splitString := strings.Split(s, sep)

	for _, entry := range splitString {

		// Some reason we have some empty entries that need skipping
		if entry == "" {
			continue
		}

		i, err := strconv.Atoi(strings.TrimSpace(entry))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		numbers = append(numbers, i)
	}

	return numbers
}
