package main

import (
	"fmt"
	"strings"

	"bufio"
	"os"
	"strconv"
	"sync"
)

func readNumbersFromFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		for _, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			numbers = append(numbers, num)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return numbers
}

func writeNumbersToFile(filename string, numbers []int) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	for _, num := range numbers {
		_, err := file.WriteString(strconv.Itoa(num) + "\n")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var numbers1, numbers2 []int
	go func() {
		defer wg.Done()
		numbers1 = readNumbersFromFile("file1.txt")
	}()

	go func() {
		defer wg.Done()
		numbers2 = readNumbersFromFile("file2.txt")
	}()

	wg.Wait()

	result := make([]int, len(numbers1))
	for i := 0; i < len(numbers1); i++ {
		result[i] = numbers1[i] * numbers2[i]
	}

	writeNumbersToFile("file3.txt", result)
}
