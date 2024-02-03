package main

import (
	"my_resume/types"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/resume", getResume)
	router.Run()
}
func getResume(c *gin.Context) {
	var resume types.Resume

	yamlData, err := os.ReadFile("resume.yml")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = yaml.Unmarshal(yamlData, &resume)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, resume)
}
