package main

import (
	"bufio"
	"fmt"
	"os"
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

	h := map[string]int64{}
	for _, str := range b {
		h[str] += 1
	}

	var ans int64
	for _, str := range a {
		i, _ := strconv.ParseInt(str, 10, 64)

		ans += h[str] * i
	}

	fmt.Println(int(ans))
}
