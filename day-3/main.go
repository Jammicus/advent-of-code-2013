package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	helpers "day3/helpers"
)

// exclude the full stop
const unicodeExclusion int = 46

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

	partOne(matrix, matrixDepth)
}

func partOne(matrix [][]rune, matrixDepth int) {

	sum := 0
	// Iterate across then down
	for i, row := range matrix {

		// Lookahead pointer
		lookAheadPointer := 0
		for j, val := range row {
			cordsToSearch := []helpers.Cordinates{}
			shoudSum := false
			if j < lookAheadPointer {
				continue
			}

			lookAheadPointer = j

			startWordCords := helpers.Cordinates{}
			endWordCords := helpers.Cordinates{}
			if unicode.IsNumber(val) {

				// Find the last Element that is a letter
				for lookAheadPointer < len(row) && unicode.IsNumber(matrix[i][lookAheadPointer]) {
					lookAheadPointer++
				}

				startWordCords = helpers.Cordinates{
					X: j,
					Y: i,
				}
				endWordCords = helpers.Cordinates{
					// As the pointer stops at the first occurance of a non number, we need to move it back 1
					// to the last occurance
					X: lookAheadPointer - 1,
					Y: i,
				}

				cordsToSearch = helpers.CalculatePerimterCordinates(startWordCords, endWordCords)

			}

			for i, cord := range cordsToSearch {
				fmt.Println(i)

				// We know we should sum, so we will
				if shoudSum {
					break
				}

				// Prevent going out of bounds on either side of the matrix
				if cord.X < 0 || cord.Y < 0 || cord.Y >= matrixDepth || cord.X >= len(row) {
					continue
				}

				fmt.Println(cord)
				if !unicode.IsDigit(matrix[cord.Y][cord.X]) && int(matrix[cord.Y][cord.X]) != unicodeExclusion {
					shoudSum = true
				}

			}

			if shoudSum {

				fmt.Println(row[startWordCords.X:endWordCords.X])
				sum = sum + helpers.ExtractWordValueFromCoords(row[startWordCords.X:endWordCords.X+1])
				shoudSum = false
			}

		}
	}

	fmt.Println(sum)
}

// func partTwo(matrix [][]rune, matrixDepth int) {

// 	sum := 0
// 	// Iterate across then down
// 	for i, row := range matrix {
// 		for j, val := range row {
// 			cordsToSearch := []helpers.Cordinates{}

// 			if string(val) == "*" {

// 				wordCords := helpers.Cordinates{
// 					X: j,
// 					Y: i,
// 				}
// 				cordsToSearch = helpers.CalculatePerimterCordinates(wordCords, wordCords)
// 			}
// 			//
// 			for i, cord := range cordsToSearch {

// 				if cord.X < 0 || cord.Y < 0 || cord.Y >= matrixDepth || cord.X >= len(row) {
// 					continue
// 				}
// 				// Fan out and find the the end of the coord
// 				if unicode.IsDigit(matrix[cord.Y][cord.X]) {
// 					leftPointer := cord.X
// 					rightPointer := cord.Y

// 				}
// 			}
// 		}
// 	}

// }
