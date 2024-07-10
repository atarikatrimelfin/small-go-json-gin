package numbercontroller

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var numbers = []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}

func SortNumbers(c *gin.Context) {
	sort.Ints(numbers)

	uniqueNumbers := make([]int, 0, len(numbers))
	seen := make(map[int]struct{})

	for _, num := range numbers {
		if _, found := seen[num]; !found {
			uniqueNumbers = append(uniqueNumbers, num)
			seen[num] = struct{}{}
		}
	}

	var uniqueNumbersStr []string
	for _, num := range uniqueNumbers {
		uniqueNumbersStr = append(uniqueNumbersStr, strconv.Itoa(num))
	}

	c.JSON(http.StatusOK, gin.H{"sorted_unique_numbers": uniqueNumbersStr})
}

func ShowDuplicates(c *gin.Context) {
	sort.Ints(numbers)

	var result []string
	currentNumber := numbers[0]
	count := 1

	for i := 1; i < len(numbers); i++ {
		if numbers[i] == currentNumber {
			count++
		} else {
			result = append(result, strconv.Itoa(currentNumber)+"["+strconv.Itoa(count)+"]")

			count = 1
		}
	}

	result = append(result, strconv.Itoa(currentNumber)+"["+strconv.Itoa(count)+"]")

	c.JSON(http.StatusOK, gin.H{"Result Duplicates Count": result})
}

func RemoveValues(c *gin.Context) {
	input := c.Query("input")
	if input == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input parameter is required"})
		return
	}

	var valuesToRemove []int
	for _, s := range strings.Split(input, ",") {
		num, err := strconv.Atoi(s)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
			return
		}
		valuesToRemove = append(valuesToRemove, num)
	}

	var result []int
	for _, num := range numbers {
		remove := false
		for _, toRemove := range valuesToRemove {
			if num == toRemove {
				remove = true
				break
			}
		}
		if !remove {
			result = append(result, num)
		}
	}

	c.JSON(http.StatusOK, gin.H{"Result": result})
}

func SumNumbersWithLimit(c *gin.Context) {
	inputStr := c.Query("input")
	input, err := strconv.Atoi(inputStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var updatedNumbers []int
	for _, num := range numbers {
		if num+input <= 10 {
			updatedNumbers = append(updatedNumbers, num+input)
		} else {
			updatedNumbers = append(updatedNumbers, 10)
		}
	}

	var updatedNumbersStr []string
	for _, num := range updatedNumbers {
		updatedNumbersStr = append(updatedNumbersStr, strconv.Itoa(num))
	}

	c.JSON(http.StatusOK, gin.H{"Result": updatedNumbersStr})
}
