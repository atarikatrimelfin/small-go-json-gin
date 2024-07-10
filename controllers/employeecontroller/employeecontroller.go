package employeecontroller

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/atarikatrimelfin/go-json-gin/models"
	"github.com/gin-gonic/gin"
)

func Smith(c *gin.Context) {
	var employee []models.Employee

	models.DB.Where("LastName LIKE ?", "Smith%").Where("TerminationDate IS NULL").Order("LastName, FirstName").Find(&employee)

	c.JSON(http.StatusOK, gin.H{"Employees who are still working with last name Smith": employee})
}

func Neverreviewed(c *gin.Context) {
	var employee []models.Employee

	models.DB.Table("employees").
		Select("employees.FirstName, employees.LastName").
		Joins("LEFT JOIN annualreviews ON employees.ID = annualreviews.empID").
		Where("annualreviews.empID IS NULL").
		Order("employees.HireDate").
		Find(&employee)

	c.JSON(http.StatusOK, gin.H{"Employees never reviewed": employee})
}

func Daydifference(c *gin.Context) {
	var firstHireDate, lastHireDate time.Time

	err := models.DB.Table("employees").
		Select("MIN(HireDate) as first_hire_date, MAX(HireDate) as last_hire_date").
		Row().
		Scan(firstHireDate, lastHireDate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employment dates"})
		return
	}

	// rumus selisih hari
	employmentDays := lastHireDate.Sub(firstHireDate).Hours() / 24

	c.JSON(http.StatusOK, gin.H{
		"first_hire_date": firstHireDate.Format("2006-01-02"),
		"last_hire_date":  lastHireDate.Format("2006-01-02"),
		"employment_days": employmentDays,
	})
}

func Salincreases(c *gin.Context) {
	var employees []struct {
		models.Employee
		TotalReviews int
	}

	err := models.DB.Table("employees").
		Select("employees.*, COUNT(annualreviews.ID) as total_reviews").
		Joins("LEFT JOIN annualreviews ON employees.ID = annualreviews.empID").
		Where("employees.TerminationDate IS NULL OR employees.TerminationDate = ?", time.Time{}).
		Group("employees.ID").
		Find(&employees).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees and reviews"})
		return
	}

	// kenaikan gaji
	for i := range employees {
		employee := &employees[i]
		initialSalary := float64(employee.Salary)
		for year := 2009; year <= 2016; year++ {
			initialSalary *= 1.15
		}
		employee.Salary = int(initialSalary)
	}

	// sortir employees
	models.DB.Table("employees").
		Select("employees.*, COUNT(annualreviews.ID) as total_reviews").
		Joins("LEFT JOIN annualreviews ON employees.ID = annualreviews.empID").
		Where("employees.TerminationDate IS NULL OR employees.TerminationDate = ?", time.Time{}).
		Group("employees.ID").
		Order("employees.Salary DESC, total_reviews ASC").
		Find(&employees)

	c.JSON(http.StatusOK, gin.H{"employees": employees})

}

func DownloadText2(c *gin.Context) {
	var employees []models.Employee

	models.DB.Where("LastName LIKE ?", "Smith%").Where("TerminationDate IS NULL").Order("LastName, FirstName").Find(&employees)

	var sb strings.Builder
	sb.WriteString("Employees who are still working with last name Smith:\n")
	for _, emp := range employees {
		sb.WriteString(fmt.Sprintf("- %s, %s\n", emp.LastName, emp.FirstName))
	}
	result := sb.String()

	c.Header("Content-Disposition", "attachment; filename=contoh2.txt")
	c.Data(http.StatusOK, "text/plain", []byte(result))
}

func DownloadText3(c *gin.Context) {
	var employees []models.Employee

	models.DB.Table("employees").
		Select("employees.FirstName, employees.LastName").
		Joins("LEFT JOIN annualreviews ON employees.ID = annualreviews.empID").
		Where("annualreviews.empID IS NULL").
		Order("employees.HireDate").
		Find(&employees)

	var sb strings.Builder
	sb.WriteString("Employees never reviewed:\n")
	for _, emp := range employees {
		sb.WriteString(fmt.Sprintf("- %s, %s\n", emp.LastName, emp.FirstName))
	}
	result := sb.String()

	c.Header("Content-Disposition", "attachment; filename=contoh3.txt")
	c.Data(http.StatusOK, "text/plain", []byte(result))
}

func DownloadText4(c *gin.Context) {
	var firstHireDate, lastHireDate time.Time

	err := models.DB.Table("employees").
		Select("MIN(HireDate) as first_hire_date, MAX(HireDate) as last_hire_date").
		Row().
		Scan(&firstHireDate, &lastHireDate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employment dates"})
		return
	}

	employmentDays := int(lastHireDate.Sub(firstHireDate).Hours() / 24)

	result := fmt.Sprintf("First Hire Date: %s\n", firstHireDate.Format("2006-01-02"))
	result += fmt.Sprintf("Last Hire Date: %s\n", lastHireDate.Format("2006-01-02"))
	result += fmt.Sprintf("Employment days: %d\n", employmentDays)

	c.Header("Content-Disposition", "attachment; filename=contoh4.txt")
	c.Data(http.StatusOK, "text/plain", []byte(result))
}

func DownloadText5(c *gin.Context) {
	var employees []struct {
		models.Employee
		TotalReviews int
	}

	err := models.DB.Table("employees").
		Select("employees.*, COUNT(annualreviews.ID) as total_reviews").
		Joins("LEFT JOIN annualreviews ON employees.ID = annualreviews.empID").
		Where("employees.TerminationDate IS NULL OR employees.TerminationDate = ?", time.Time{}).
		Group("employees.ID").
		Find(&employees).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees and reviews"})
		return
	}

	for i := range employees {
		employee := &employees[i]
		initialSalary := float64(employee.Salary)
		for year := 2009; year <= 2016; year++ {
			initialSalary *= 1.15
		}
		employee.Salary = int(initialSalary)
	}

	err = models.DB.Table("employees").
		Select("employees.*, COUNT(annualreviews.ID) as total_reviews").
		Joins("LEFT JOIN annualreviews ON employees.ID = annualreviews.empID").
		Where("employees.TerminationDate IS NULL OR employees.TerminationDate = ?", time.Time{}).
		Group("employees.ID").
		Order("employees.Salary DESC, total_reviews ASC").
		Find(&employees).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sorted employees and reviews"})
		return
	}

	var sb strings.Builder
	sb.WriteString("Employees sorted by salary increase and total reviews:\n")
	for _, emp := range employees {
		sb.WriteString(fmt.Sprintf("- %s, %s: Salary %d, Total Reviews %d\n", emp.LastName, emp.FirstName, emp.Salary, emp.TotalReviews))
	}
	result := sb.String()

	c.Header("Content-Disposition", "attachment; filename=salincreases.txt")
	c.Data(http.StatusOK, "text/plain", []byte(result))
}

// SHOW FILE
func ShowText2(c *gin.Context) {
	filename := "contoh2.txt"

	file, err := os.Open(filename)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file info"})
		return
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	c.Data(http.StatusOK, "text/plain", buffer)
}

func ShowText3(c *gin.Context) {
	filename := "contoh3.txt"

	file, err := os.Open(filename)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file info"})
		return
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	c.Data(http.StatusOK, "text/plain", buffer)
}

func ShowText4(c *gin.Context) {
	filename := "contoh4.txt"

	file, err := os.Open(filename)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file info"})
		return
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	c.Data(http.StatusOK, "text/plain", buffer)
}

func ShowText5(c *gin.Context) {
	filename := "contoh5.txt"

	file, err := os.Open(filename)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file info"})
		return
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	c.Data(http.StatusOK, "text/plain", buffer)
}
