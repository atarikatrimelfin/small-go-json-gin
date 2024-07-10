package main

import (
	"github.com/atarikatrimelfin/go-json-gin/controllers/citycontroller"
	"github.com/atarikatrimelfin/go-json-gin/controllers/employeecontroller"
	"github.com/atarikatrimelfin/go-json-gin/controllers/numbercontroller"
	"github.com/atarikatrimelfin/go-json-gin/controllers/rantcontroller"
	"github.com/atarikatrimelfin/go-json-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/json/smith", employeecontroller.Smith)
	r.GET("/json/neverreviewed", employeecontroller.Neverreviewed)
	r.GET("/json/daydifference", employeecontroller.Daydifference)
	r.GET("/json/salincreases", employeecontroller.Salincreases)

	r.GET("/download/smith", employeecontroller.DownloadText2)
	r.GET("/download/neverreviewed", employeecontroller.DownloadText3)
	r.GET("/download/daydifference", employeecontroller.DownloadText4)
	r.GET("/download/salincreases", employeecontroller.DownloadText5)

	r.GET("/showfile/smith", employeecontroller.ShowText2)
	r.GET("/showfile/neverreviewed", employeecontroller.ShowText3)
	r.GET("/showfile/daydifference", employeecontroller.ShowText4)
	r.GET("/showfile/salincreases", employeecontroller.ShowText5)

	r.GET("/result8a", citycontroller.CheckCity)

	r.GET("/sortnumbers", numbercontroller.SortNumbers)
	r.GET("/showduplicates", numbercontroller.ShowDuplicates)
	r.GET("/removevalues", numbercontroller.RemoveValues)
	r.GET("/sumnumber", numbercontroller.SumNumbersWithLimit)

	r.GET("/random", rantcontroller.GenerateRandomString)
	// r.GET("/random2", rantcontroller.SortRandomString)
	// r.GET("/random3", rantcontroller.SortRandomString2)

	r.Run()
}
