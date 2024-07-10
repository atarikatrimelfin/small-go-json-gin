package citycontroller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var cities = []string{"Bandung", "Cimahi", "Ambon", "Jayapura", "Makasar"}

func CheckCity(c *gin.Context) {
	cityName := c.Query("city")

	found := false
	for _, city := range cities {
		if city == cityName {
			found = true
			break
		}
	}

	if found {
		c.JSON(http.StatusOK, gin.H{"result": true})
	} else {
		suggestions := getSuggestions(cityName)
		c.JSON(http.StatusOK, gin.H{"Result": false, "City Suggestions": suggestions})
	}
}

func getSuggestions(city string) string {
	var suggestions []string

	firstChar := strings.ToUpper(city[:1])
	for _, c := range cities {
		if strings.HasPrefix(strings.ToUpper(c), firstChar) {
			suggestions = append(suggestions, c)
		}
	}

	lastChar := strings.ToUpper(city[len(city)-1:])
	for _, c := range cities {
		if strings.HasSuffix(strings.ToUpper(c), lastChar) && !contains(suggestions, c) {
			suggestions = append(suggestions, c)
		}
	}

	suggestionsStr := strings.Join(suggestions, ", ")

	return suggestionsStr
}

func contains(arr []string, city string) bool {
	for _, c := range arr {
		if c == city {
			return true
		}
	}
	return false
}
