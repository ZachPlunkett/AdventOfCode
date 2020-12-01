package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// expenseReport will determine if 2 entries in a given slice add up to
// 2020. If so their product will be returned
func expenseReport(entries []int) int {
	for i, entry1 := range entries {
		for _, entry2 := range entries[i:] {
			if entry1+entry2 == 2020 {
				return entry1 * entry2
			}
		}
	}
	return -1
}

func expenseReportThree(entries []int) int {
	for i, entry1 := range entries {
		for j, entry2 := range entries[i:] {
			for _, entry3 := range entries[j:] {
				if entry1+entry2+entry3 == 2020 {
					fmt.Println(entry1, entry2, entry3)
					return entry1 * entry2 * entry3
				}
			}
		}
	}
	return -1
}

func readFile(fileName string) ([]int, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	intStrings := strings.Split(string(fileBytes), "\n")

	// Set size initally to avoid unnecessary resizes
	ints := make([]int, 0, len(intStrings))

	for _, entry := range intStrings {
		entry = strings.TrimSpace(entry)
		if len(entry) == 0 {
			continue
		}

		n, err := strconv.Atoi(entry)
		if err != nil {
			return nil, err
		}

		ints = append(ints, n)
	}
	return ints, nil
}

func main() {
	entries, err := readFile("input.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println(expenseReport(entries))
	fmt.Println(expenseReportThree(entries))
}
