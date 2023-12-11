package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type node struct {
	label string
	left  *node
	right *node
}

func main() {
	lineCounter := 0
	fileLocation := os.Getenv("DAY8_FILE")
	instructions := ""
	nodeList := []*node{}

	partOneStartNode := node{}
	partTwoStartNodes := []*node{}

	nodeMap := map[string][]string{}

	if fileLocation == "" {
		fallback := "day8.txt"
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

		if lineCounter == 0 {
			instructions = scanner.Text()
			lineCounter++
			continue
		}

		if lineCounter == 1 {
			lineCounter++
			continue
		}

		s := strings.Split(scanner.Text(), "=")

		lr := strings.Split(s[1], ",")

		name := strings.TrimSpace(s[0])
		left := strings.TrimSpace(strings.TrimPrefix(lr[0], " ("))
		right := strings.TrimSpace(strings.TrimSuffix(lr[1], ")"))

		nodeMap[name] = []string{left, right}

	}

	// create our node list
	for k, _ := range nodeMap {
		n := &node{
			label: k,
		}

		nodeList = append(nodeList, n)
	}

	for i, node := range nodeList {

		n := nodeMap[node.label]
		l := n[0]
		r := n[1]

		for _, node1 := range nodeList {

			if node1.label == l {
				node.left = node1
			}

			if node1.label == r {
				node.right = node1
			}

		}

		if node.label == "AAA" {
			partOneStartNode = *node
			fmt.Println(node.label)
		}

		if strings.HasSuffix(node.label, "A") {
			partTwoStartNodes = append(partTwoStartNodes, node)
		}

		nodeList[i] = node

	}

	fmt.Println(partOneStartNode)
	fmt.Printf("part one: %v \n", partOne(instructions, partOneStartNode))
	fmt.Printf("part two: %v \n", partTwo(instructions, partTwoStartNodes))

}

func partOne(instructions string, startNode node) int {

	counter := 0
	i := 0
	node := startNode

	for node.label != "ZZZ" {

		fmt.Println("moving through node ", node.label)

		counter++

		if i == len(instructions) {
			i = 0
		}

		if string(instructions[i]) == "L" {

			node = *node.left
			i++
			continue

		}

		node = *node.right
		i++

	}

	return counter
}

func partTwo(instructions string, startNodes []*node) int {

	lcm := []int{}

	for _, node := range startNodes {
		fmt.Printf("%v, left: %v, right: %v\n", node.label, node.left.label, node.right.label)
		counter := 0
		i := 0
		for !strings.HasSuffix(node.label, "Z") {
			if i == len(instructions) {
				i = 0
			}

			instruction := instructions[i]

			if i == len(instructions) {
				i = 0
			}

			if instruction == 76 {
				node = node.left
			}

			if instruction == 82 {
				node = node.right
			}
			counter++
			i++

		}

		lcm = append(lcm, counter)

	}

	return findLCM(lcm)
}

func findLCM(vals []int) int {
	hcm := findHCM(vals)
	pFactors := primeFactors(hcm, vals)

	x := 1
	for _, pf := range pFactors {
		x = x * pf
	}

	x = x * hcm

	return x
}

func findHCM(vals []int) int {

	fmt.Println("Finding HCM")
	slices.Sort(vals)

	i := vals[len(vals)-1]
	for {
		counter := 0
		for _, val := range vals {

			if val%i == 0 {
				counter++
			}
		}

		if counter == len(vals) {
			fmt.Println("HCM: ", counter)
			return i
		}
		i--
	}

}

func primeFactors(hcm int, vals []int) []int {
	fmt.Println("Finding prime factors")
	pFactors := []int{}
	for _, val := range vals {

		for i := 0; i < val; i++ {

			if hcm*i == val {
				fmt.Println("Found PF for val ", val, i)
				pFactors = append(pFactors, i)
				continue
			}

		}
	}
	return pFactors
}
