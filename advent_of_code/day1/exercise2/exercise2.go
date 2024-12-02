package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"strings"
)
// Calculate a total similarity score by adding up each number in the left list after multiplying it by the number of times that number appears in the right list.
// Input will be given via stdin in the format 0 1\n2 4\n ...
// to do this, we want to parse all the left and right numbers
// It's important to consider that maybe a number is shown in the right before it's shown in the left, so we should do this after parsing all the numbers
// Essentially we are going to want to create a frequency map out of rightNumbers to avoid doing O(N^2) by looping over RightVals for each number in LeftVals
// This involves building a map of type int[int]

func createFrequencyMap(nums []int) (m map[int]int) {
	m = make(map[int]int)
	for _, v := range nums {
		m[v]++ 
	}
	return
}

func main() {
	// takes O(N) time and space
	// set up our scanner line by line
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	// initialize our arrays as empty
	var leftNumbers, rightNumbers []int

	// begin scanning lines and parsing into values for the left and right list
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		left, _ := strconv.Atoi(words[0])
		right, _ := strconv.Atoi(words[1])
		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)
	}

	// create frequency map
	freq := createFrequencyMap(rightNumbers)

	// total the similarity score for all numbers in leftNumbers
	total := 0
	for _, v := range leftNumbers {
		if x, ok := freq[v]; ok {
			total = total + v * x
		}
	}
	fmt.Println(total)
}
