package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Cordinates struct {
	X int
	Y int
}

// exclude the full stop
const unicodeExclusion int = 46

//TODO: Part 2

func main() {
	fileLocation := os.Getenv("DAY3_FILE")
	matrixDepth := 0
	matrix := make([][]rune, 0)
	sum := 0

	if fileLocation == "" {
		fallback := "day-3/day3.txt"
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
		entry := stringToArray(scanner.Text())
		tmp := make([]rune, 0)
		tmp = append(tmp, entry...)
		matrix = append(matrix, tmp)
		matrixDepth++
	}

	// Iterate across then down
	for i, row := range matrix {

		// Lookahead pointer
		lookAheadPointer := 0
		for j, val := range row {
			cordsToSearch := []Cordinates{}
			shoudSum := false
			if j < lookAheadPointer {
				continue
			}

			lookAheadPointer = j

			startWordCords := Cordinates{}
			endWordCords := Cordinates{}
			if unicode.IsNumber(val) {

				// Find the last Element that is a letter
				for lookAheadPointer < len(row) && unicode.IsNumber(matrix[i][lookAheadPointer]) {
					lookAheadPointer++
				}

				startWordCords = Cordinates{
					X: j,
					Y: i,
				}
				endWordCords = Cordinates{
					// As the pointer stops at the first occurance of a non number, we need to move it back 1
					// to the last occurance
					X: lookAheadPointer - 1,
					Y: i,
				}

				cordsToSearch = calculatePerimterCordinates(startWordCords, endWordCords)

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
				sum = sum + extractWordValueFromCoords(row[startWordCords.X:endWordCords.X+1])
				shoudSum = false
			}

		}
	}

	fmt.Println(sum)
}

func extractWordValueFromCoords(matrix []rune) int {

	str := ""
	for _, v := range matrix {
		fmt.Println(string(v))
		str = fmt.Sprintf("%v%v", str, string(v))
	}

	fmt.Println(str)

	val, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err)
		return -100000
	}

	fmt.Println("Found: ", val)

	return val
}

func stringToArray(input string) []rune {
	runes := make([]rune, len(input))

	for j, char := range input {

		runes[j] = char

	}

	return runes
}

// We assume its always left to right, not diagonal
func calculatePerimterCordinates(startIndex, endIndex Cordinates) []Cordinates {
	combinedCordinates := []Cordinates{}
	leftMostCords := []Cordinates{}
	rightMostCords := []Cordinates{}

	//Calculate the left and right most cords.
	// Index 1 will always be the directly left or right.
	for i := -1; i < 2; i++ {
		lCoord := Cordinates{
			X: startIndex.X - 1,
			Y: startIndex.Y - i,
		}

		rCoord := Cordinates{
			X: endIndex.X + 1,
			Y: startIndex.Y - i,
		}

		leftMostCords = append(leftMostCords, lCoord)
		rightMostCords = append(rightMostCords, rCoord)

	}

	// Above the line
	for i := leftMostCords[1].X; i <= rightMostCords[1].X; i++ {
		c := Cordinates{
			X: i,
			Y: startIndex.Y - 1,
		}

		combinedCordinates = append(combinedCordinates, c)
	}

	// Below the line
	for i := leftMostCords[1].X; i <= rightMostCords[1].X; i++ {
		c := Cordinates{
			X: i,
			Y: startIndex.Y + 1,
		}

		combinedCordinates = append(combinedCordinates, c)
	}

	combinedCordinates = append(combinedCordinates, leftMostCords...)
	combinedCordinates = append(combinedCordinates, rightMostCords...)

	return combinedCordinates
}
