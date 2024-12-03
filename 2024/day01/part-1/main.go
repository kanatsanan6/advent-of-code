package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("2024/day-1/input.txt")
	defer file.Close()

	var a, b []string

	s := bufio.NewScanner(file)
	for s.Scan() {
		data := strings.Fields(s.Text())

		a = append(a, data[0])
		b = append(b, data[1])
	}

	slices.Sort(a)
	slices.Sort(b)

	var ans float64
	for i := 0; i < len(a); i++ {
		x, _ := strconv.ParseFloat(a[i], 64)
		y, _ := strconv.ParseFloat(b[i], 64)

		ans += math.Abs(y - x)
	}

	fmt.Println(int(ans))
}
