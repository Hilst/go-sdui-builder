package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"sdui.builder/builder"
	"sdui.builder/models"
)

// MAIN
func main() {
	router := gin.Default()
	router.GET("/", getIndex)
	router.GET("/sample", getSample)
	router.POST("/build", postBuild)

	router.Run("localhost:8080")
}

// GET INDEX
func getIndex(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// GET SAMPLE
func getSample(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "SDUI Builder is here in GO")
}

// POST BUILD
func postBuild(c *gin.Context) {
	var input models.Input

	if err := c.BindJSON(&input); err != nil {
		fmt.Println("\n" + err.Error() + "\n")
		return
	}

	var dto models.LayoutDTO = builder.BuildLayoutDTO(input.Data, input.Layout)
	c.IndentedJSON(http.StatusOK, dto)
}
