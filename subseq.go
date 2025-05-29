package main

import (
	"bufio"
	"fmt"
	"os"
)

var in *bufio.Reader
var out *bufio.Writer

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	posled := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &posled[i])
	}
	if n < 26 {
		fmt.Fprintln(out, "NONE")
		return
	}

	result := solve(n, posled)
	if result == -1 {
		fmt.Fprintln(out, "NONE")
	} else {
		fmt.Fprintln(out, result)
	}
}

func solve(n int, posled []int) int {
	count := make([]int, 27)
	diff := 0
	left := 0
	minLen := n + 1

	for right := 0; right < n; right++ {
		val := posled[right]
		if val > 0 && val < 27 {
			if count[val] == 0 {
				diff++
			}
			count[val]++
		}
		for diff == 26 {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			lval := posled[left]
			if lval > 0 && lval < 27 {
				count[lval]--
				if count[lval] == 0 {
					diff--
				}
			}
			left++

		}
	}
	if minLen == n+1 {
		return -1
	}
	return minLen
}
