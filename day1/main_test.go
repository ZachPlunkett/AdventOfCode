package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpenseReport(t *testing.T) {
	assert := assert.New(t)

	// Test using correct example given on AOC
	given := []int{1721, 979, 366, 299, 675, 1456}
	answer := 514579
	output := expenseReport(given)
	assert.Equal(answer, output, "Should match given sample output")

	// Test to make sure output is correct for when no answer is found
	output = expenseReport([]int{1, 2, 3, 4, 5})
	assert.Equal(output, -1, "Return should be -1 when no answer is found")
}

func TestExpenseReportThree(t *testing.T) {
	assert := assert.New(t)

	// Test using correct example given on AOC
	given := []int{1721, 979, 366, 299, 675, 1456}
	answer := 241861950
	output := expenseReportThree(given)
	assert.Equal(answer, output, "Should match given sample output")

	// Test to make sure output is correct for when no answer is found
	output = expenseReportThree([]int{1, 2, 3, 4, 5})
	assert.Equal(output, -1, "Return should be -1 when no answer is found")
}
