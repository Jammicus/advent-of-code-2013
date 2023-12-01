package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	numberMap = map[string]int{
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
)

func main() {
	sum := 0
	fileLocation := os.Getenv("DAY1_FILE")

	if fileLocation == "" {
		fallback := "day-1/day1.txt"
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

		number, err := lineHustler(scanner.Text())

		if err != nil {
			fmt.Println("Error parsing numbers in line: %v", scanner.Text())
		}
		sum = sum + number
		fmt.Println("Line %v gave us %v", scanner.Text(), number)
	}

	fmt.Println("total: %v", sum)
}

// Hustle through the line from both directions. Got no times to iterate twice
func lineHustler(line string) (int, error) {

	// Valid numbers 0-9, use negatives to represent invalid numbers
	leftHandNum := -1
	rightHandNum := -1
	leftHandIndexFound := -1
	rightHandIndexFound := -1
	length := len(line) - 1

	for i := 0; i < length; i++ {
		lhIndex := i
		rhIndex := length - i

		if leftHandNum != -1 && rightHandNum != -1 {
			break
		}

		// Left to right
		if val := letterToNumber(string(line[lhIndex])); val != -1 && leftHandNum == -1 {

			// Traversed through to the found right item
			if lhIndex == rightHandIndexFound {
				break
			}
			leftHandNum = val
			leftHandIndexFound = lhIndex
		}

		if val := wordToNumber(string(line[lhIndex:]), true); val != -1 && leftHandNum == -1 {

			// Traversed through to the found right item
			if lhIndex == rightHandIndexFound {
				break
			}
			leftHandNum = val
			leftHandIndexFound = lhIndex
		}

		//Right to left
		if val := letterToNumber(string(line[rhIndex])); val != -1 && rightHandNum == -1 {

			//Travered through and hit left hand item
			if rhIndex == leftHandIndexFound {
				break
			}

			rightHandNum = val
			rightHandIndexFound = rhIndex
		}

		if val := wordToNumber(string(line[:rhIndex+1]), false); val != -1 && rightHandNum == -1 {

			// Traversed through to the found right item
			if lhIndex == rightHandIndexFound {
				break
			}
			rightHandNum = val
			rightHandIndexFound = rhIndex
		}
	}

	fmt.Println(leftHandNum, rightHandNum)

	if leftHandNum == -1 && rightHandNum != -1 {
		return strconv.Atoi(fmt.Sprintf("%d%d", rightHandNum, rightHandNum))
	}

	if leftHandNum != -1 && rightHandNum == -1 {
		return strconv.Atoi(fmt.Sprintf("%d%d", leftHandNum, leftHandNum))
	}

	return strconv.Atoi(fmt.Sprintf("%d%d", leftHandNum, rightHandNum))
}

// LTR returns true if the number was found left -> right in the string. False if right -> left
func wordToNumber(s string, ltr bool) (number int) {

	for key, val := range numberMap {
		if ltr && strings.HasPrefix(s, key) {
			return val
		}
		if !ltr && strings.HasSuffix(s, key) {
			return val
		}
	}
	return -1
}

// Return -1 if not a number
func letterToNumber(s string) int {
	val, err := strconv.Atoi(s)

	if err != nil {
		return -1
	}
	return val
}
