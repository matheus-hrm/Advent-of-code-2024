package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read(path string) ([]int, []int, error) {
	var col1, col2 []int

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		col := strings.Fields(line)

		x, err1 := strconv.Atoi(col[0])
		y, err2 := strconv.Atoi(col[1])
		if err1 != nil || err2 != nil {
			log.Printf("Skipping line with invalid numbers: %q", line)
			continue
		}
		col1 = append(col1, x)
		col2 = append(col2, y)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return col1, col2, nil
}

func main() {
	x, y, err := read("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	sort.Ints(x)
	sort.Ints(y)

	var sum float64
	freq_y := make(map[int]int)
	for i := 0; i < len(x); i++ {
		d := math.Abs(float64(x[i]) - float64(y[i]))
		freq_y[y[i]]++
		sum += d
	}

	similarity := 0
	for i := 0; i < len(x); i++ {
		if freq_y[x[i]] > 0 {
			similarity += freq_y[x[i]] * x[i]
		}
	}
	log.Printf("Result: %.f", sum)
	log.Printf("Similarity: %d", similarity)
}
