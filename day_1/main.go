package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ans [3]int
	curr := 0
	end := false
	for {
		scanner.Scan()
		line := scanner.Text()

		if len(line) == 0 {
			if end {
				break
			}
			end = true
			curr = 0
			continue
		}

		end = false
		lineVal, _ := strconv.Atoi(line)
		curr += lineVal
		if curr > ans[0] {
			ans[1] = ans[0]
			ans[0] = curr
			ans[2] = ans[1]
			continue
		}

		if curr > ans[1] {
			ans[2] = ans[1]
			ans[1] = curr
			continue
		}

		if curr > ans[2] {
			ans[2] = curr
		}
	}

	fmt.Printf("Answer: %d", ans[0]+ans[1]+ans[2])
}
