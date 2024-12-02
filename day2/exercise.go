package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"strings"
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func validOrder(l1, l2 int, inc bool) bool {
	diff := l2 - l1
	if diff < 0 && inc {
		return false
	} 
	if diff > 0 && !inc {
		return false
	}
	return true
}

func validDiff(l1, l2 int) bool {
	diff := abs(l2 - l1)
	return diff >= 1 && diff <= 3
}

func validReport(nums []int) (res bool) {
	// adjacent levels must differ by at least 1 and at most 3
	// levels must be either all increasing or all decreasing

	// if the length of levels is 0 or 1 just say the report passed
	if len(nums) <= 1 {
		return true
	}

	// set up an initial bool result which if contradicted will result in us returning false during the loop
	inc := !(nums[len(nums) - 1] - nums[0] < 0)

	pre := nums[0]

	// do the main loop 
	for _, x := range nums[1:] {
		if !validOrder(pre, x, inc) || !validDiff(pre, x) {
			return false
		}
		pre = x
	}
	return true
}

func main() {
	// takes O(NlogN) time, O(N) space
	// Pair up the numbers and measure how far apart they are. The numbers will be given via stdin i.e. 3 4\n 4 5\n ...
	// Pair up the smallest number in the left list with the smallest number in the right list, then the second-smallest left number with the second-smallest right number, and so on.

	// set up our scanner line by line
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	successfulReports := 0

	// begin scanning lines and parsing into values for the left and right list
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		levels := make([]int, len(words))

		for i, x := range words {
			levels[i], _ = strconv.Atoi(x)
		}

		if validReport(levels) {
			successfulReports = successfulReports + 1
		}	
	}

	fmt.Println(successfulReports)

	os.Exit(0)
}
