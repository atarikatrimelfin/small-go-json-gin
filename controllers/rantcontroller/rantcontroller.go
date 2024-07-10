package rantcontroller

import (
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type RandomStringResult struct {
	RandomString    string `json:"RandomString"`
	SortedString    string `json:"SortedString"`
	TotalLetters    int    `json:"TotalLetters"`
	TotalVowels     int    `json:"TotalVowels"`
	TotalDigits     int    `json:"TotalDigits"`
	TotalEvenDigits int    `json:"TotalEvenDigits"`
}

// RANDOM VALUES GENERATE
func GenerateRandomString(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())

	var letters strings.Builder
	for i := 0; i < 50; i++ {
		letters.WriteByte(byte(rand.Intn(26)) + 'a')
	}

	var digits strings.Builder
	for i := 0; i < 50; i++ {
		digits.WriteByte(byte(rand.Intn(10)) + '0')
	}

	randomString := letters.String() + digits.String()

	sortedString := sortRandomString(randomString)

	totalLetters := countLetters(randomString)
	totalVowels := countVowels(randomString)
	totalDigits := countDigits(randomString)
	totalEvenDigits := countEvenDigits(randomString)

	result := RandomStringResult{
		RandomString:    randomString,
		SortedString:    sortedString,
		TotalLetters:    totalLetters,
		TotalVowels:     totalVowels,
		TotalDigits:     totalDigits,
		TotalEvenDigits: totalEvenDigits,
	}

	c.JSON(http.StatusOK, result)
}

func countLetters(s string) int {
	count := 0
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			count++
		}
	}
	return count
}

func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func countDigits(s string) int {
	count := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			count++
		}
	}
	return count
}

func countEvenDigits(s string) int {
	count := 0
	for _, char := range s {
		digit := int(char - '0')
		if digit >= 0 && digit <= 9 && digit%2 == 0 {
			count++
		}
	}
	return count
}

func sortRandomString(s string) string {
	var letters, digits []byte

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			letters = append(letters, s[i])
		} else if s[i] >= '0' && s[i] <= '9' {
			digits = append(digits, s[i])
		}
	}

	sort.Slice(digits, func(i, j int) bool {
		return digits[i] > digits[j]
	})

	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	sortedString := string(digits) + string(letters)

	return sortedString
}
