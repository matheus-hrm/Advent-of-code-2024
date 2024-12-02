package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func read(path string) ([][]int, error) {
	matrix := make([][]int, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		lines := scanner.Text()
		line := strings.Split(lines, "\n")

		for _, l := range line {
			if len(l) == 0 {
				continue
			}
			matrix = append(matrix, []int{})
			for _, v := range strings.Split(l, " ") {
				x, err := strconv.Atoi(v)
				if err != nil {
					log.Printf("Skipping line with invalid numbers: %q", l)
					continue
				}
				matrix[i] = append(matrix[i], x)
			}
			i++
			log.Printf("Read line: %v", matrix[i-1])
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return matrix, nil
}

func main() {
	matrix, err := read("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	safe := make([]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		if isRowSafe(matrix[i]) {
			safe[i] = 1
		}
	}

	sum := 0
	for _, v := range safe {
		sum += v
	}
	log.Printf("Safe: %d", sum)

	for i := 0; i < len(matrix); i++ {
		if isRowSafe(matrix[i]) {
			continue
		}
		for j := 0; j < len(matrix[i]); j++ {
			newRow := append([]int{}, matrix[i][:j]...)
			newRow = append(newRow, matrix[i][j+1:]...)
			log.Printf("New row: %v", newRow)
			if isRowSafe(newRow) {
				safe[i] = 1
				break
			}
		}
	}
	sum = 0
	for _, v := range safe {
		sum += v
	}
	log.Printf("Safe: %d", sum)
}

func isRowSafe(row []int) bool {
	isIncreasing := row[0] < row[1]

	for i := 0; i < len(row)-1; i++ {
		if math.Abs(float64(row[i]-row[i+1])) > 3 {
			return false
		}
		if isIncreasing {
			if row[i] >= row[i+1] {
				return false
			}
		} else {
			if row[i] <= row[i+1] {
				return false
			}
		}
	}
	return true
}
