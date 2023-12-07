package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"

	helpers "day3/helpers"
)

// exclude the full stop
const unicodeExclusion int = 46
const unicodeAsterix int = 42

//TODO: Part 2

func main() {
	fileLocation := os.Getenv("DAY3_FILE")
	matrixDepth := 0
	matrix := make([][]rune, 0)

	if fileLocation == "" {
		fallback := "day3.txt"
		fmt.Printf("Falling back to %v as james is a dummy and didnt give me a file \n", fallback)
		fileLocation = fallback
	}

	file, err := os.Open(fileLocation)

	if err != nil {
		fmt.Println("Error finding file %v", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Load up the 2d matrix
		entry := helpers.StringToArray(scanner.Text())
		tmp := make([]rune, 0)
		tmp = append(tmp, entry...)
		matrix = append(matrix, tmp)
		matrixDepth++
	}

	x, y := matrixParser(matrix)
	p1 := partOneRefactored(x, y)
	p2 := partTwoRefactored(x, y)
	fmt.Println("Part one answer ", p1)
	fmt.Println("Part Two Answer ", p2)

}

func matrixParser(matrix [][]rune) (numberCordinates []helpers.NumberCoordinates, specialCharCoordinates []helpers.Coordinates) {
	numberCordinates = []helpers.NumberCoordinates{}
	specialCharCoordinates = []helpers.Coordinates{}
	for i, row := range matrix {

		wCord := helpers.NumberCoordinates{
			StartIndexXAxis: -1,
			EndIndexXAxis:   -1,
		}

		for j, val := range row {

			// We have a open cord and need to close it
			if (wCord.StartIndexXAxis != -1 && !unicode.IsNumber(val)) || (wCord.StartIndexXAxis != -1 && j == len(row)-1) {

				if j == len(row)-1 && unicode.IsNumber(val) {

					val, err := strconv.Atoi(fmt.Sprintf("%d%s", wCord.Value, string(val)))
					if err != nil {
						log.Fatal("Parsing error ", err)
					}
					wCord.Value = val

				}
				fmt.Println(wCord.Value)
				numberCordinates = append(numberCordinates, wCord)

				wCord = helpers.NumberCoordinates{
					StartIndexXAxis: -1,
					EndIndexXAxis:   -1,
					YAxis:           -1,
					Value:           0,
				}
			}

			if !unicode.IsNumber(val) && int(val) != unicodeExclusion {
				cord := helpers.Coordinates{
					X:     j,
					Y:     i,
					Value: val,
				}

				specialCharCoordinates = append(specialCharCoordinates, cord)

			}

			if unicode.IsNumber(val) {
				if wCord.StartIndexXAxis == -1 {
					wCord.StartIndexXAxis = j
					wCord.YAxis = i
				}

				val, err := strconv.Atoi(fmt.Sprintf("%d%s", wCord.Value, string(val)))
				if err != nil {
					log.Fatal("Parsing error ", err)
				}
				wCord.Value = val
				wCord.EndIndexXAxis = j
			}

		}
	}

	fmt.Printf("Returing %v number cordinates and %v special character cordinates \n", len(numberCordinates), len(specialCharCoordinates))
	return numberCordinates, specialCharCoordinates
}

func partOneRefactored(numberCordinates []helpers.NumberCoordinates, specialCharCoordinates []helpers.Coordinates) int {

	sum := 0
	for _, specialCord := range specialCharCoordinates {

		cordsToSearch := helpers.CalculatePerimterCoordinates(specialCord)

		for _, numberCord := range numberCordinates {
			found := false
			for _, searchCord := range cordsToSearch {
				if searchCord.Y == numberCord.YAxis {
					if searchCord.X >= numberCord.StartIndexXAxis && searchCord.X <= numberCord.EndIndexXAxis && !found {
						sum = sum + numberCord.Value
						found = true
					}
				}
			}
		}

	}

	return sum

}

func partTwoRefactored(numberCordinates []helpers.NumberCoordinates, specialCharCoordinates []helpers.Coordinates) int {
	sum := 0
	for _, specialCord := range specialCharCoordinates {

		if int(specialCord.Value) != unicodeAsterix {
			continue
		}
		foundVal := 0
		foundVals := []int{}
		cordsToSearch := helpers.CalculatePerimterCoordinates(specialCord)
		for _, numberCord := range numberCordinates {
			for _, searchCord := range cordsToSearch {

				if searchCord.Y == numberCord.YAxis {

					if searchCord.X >= numberCord.StartIndexXAxis && searchCord.X <= numberCord.EndIndexXAxis && numberCord.Value != foundVal {

						foundVals = append(foundVals, numberCord.Value)
						break

					}

				}
			}
		}

		if len(foundVals) == 2 {
			sum = sum + (foundVals[0] * foundVals[1])
		}
	}

	return sum

}
