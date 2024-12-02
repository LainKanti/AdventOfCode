package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"bufio"
	"strings"
)

// set up an IntHeap
type IntHeap []int

func (h IntHeap) Len() int 	{ return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) 	{ h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// go doesn't have an abs function so we'll implement one
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {
	// takes O(NlogN) time, O(N) space
	// Pair up the numbers and measure how far apart they are. The numbers will be given via stdin i.e. 3 4\n 4 5\n ...
	// Pair up the smallest number in the left list with the smallest number in the right list, then the second-smallest left number with the second-smallest right number, and so on.

	// set up our scanner line by line
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	// initialize the total difference between the smallest elements
	total := 0

	// initialize our heaps
	lh := &IntHeap{}
	heap.Init(lh)

	rh := &IntHeap{}
	heap.Init(rh)

	// begin scanning lines and parsing into values for the left and right list
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		left, _ := strconv.Atoi(words[0])
		right, _ := strconv.Atoi(words[1])

		// push left and right values onto respective heaps
		heap.Push(lh, left)
		heap.Push(rh, right)
	}

	// loop through heaps and pop respective lowest values, get the absolute difference, add this to the total and write it to stdout
	for lh.Len() > 0 && rh.Len() > 0 {
		leftVal := heap.Pop(lh).(int)
		rightVal := heap.Pop(rh).(int)
		total = abs(leftVal - rightVal) + total
	}

	fmt.Println(total)
}
