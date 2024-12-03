package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"regexp"
)

/*
func mult(nums ...int) int {
	total := 1 
	for _, x := range nums {
		total = total * x
	}
	return total
}
*/

func getAllMultStrings(text string) []string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	return re.FindAllString(text, -1)
}
/*
func getAllIntegersFromString(text string) []int {
	re2 := regexp.MustCompile(`\d+`)
	matches2 := re2.FindAllString(text, -1)
	ints := make([]int, len(matches2))
	for i, x := range matches2 {
		ints[i], _ = strconv.Atoi(x)
	}
	return ints
}
*/
func main() {
	// takes O(NlogN) time, O(N) space
	// Pair up the numbers and measure how far apart they are. The numbers will be given via stdin i.e. 3 4\n 4 5\n ...
	// Pair up the smallest number in the left list with the smallest number in the right list, then the second-smallest left number with the second-smallest right number, and so on.

	// set up our scanner line by line
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	// begin scanning lines and parsing into values for the left and right list
	total := 0
	for scanner.Scan() {
		// doesn't support multiple lines atm 
		line := scanner.Text()
		matches := getAllMultStrings(line)
		for _, x := range matches {
			re2 := regexp.MustCompile(`\d+`)
			matches2 := re2.FindAllString(x, -1)
			total2 := 1
			for _, x := range matches2 {
				x, _ := strconv.Atoi(x)
				total2 = x * total2
			}
			total = total + total2
		}
	}
	fmt.Println(total)
	os.Exit(0)
}
