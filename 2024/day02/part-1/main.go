package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("2024/day-2/input.txt")
	defer file.Close()

	var ans int

	s := bufio.NewScanner(file)
	for s.Scan() {
		isSafe := true
		report := strings.Fields(s.Text())
		prevDiff := float64(0)

		for i := 1; i < len(report); i++ {
			x, _ := strconv.ParseFloat(report[i-1], 64)
			y, _ := strconv.ParseFloat(report[i], 64)
			diff := x - y

			if i != 1 && !((diff < 0 && prevDiff < 0) || (diff > 0 && prevDiff > 0)) || math.Abs(diff) > 3 {
				isSafe = false
				break
			}

			prevDiff = diff
		}

		if isSafe {
			ans += 1
		}
	}

	fmt.Println(ans)
}
