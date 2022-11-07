package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/theMariusK/algorithms/sorting"
)

const (
	INCR int = 1
	DECR int = 2
)

func getNumbers(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var arr []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		arr = append(arr, n)
	}

	return arr
}

func main() {
	arr := getNumbers("C:/Users/marix/Documents/Projects/algorithms/samples/numbers.txt")

	fmt.Printf("Original: %d\n", arr)
	start := time.Now()
	sorted := sorting.InsertionSort(arr, INCR)
	elapsed := time.Since(start)
	fmt.Printf("Sorted: %d\n", sorted)

	fmt.Printf("\nTime elapsed: %s\n", elapsed)
}
