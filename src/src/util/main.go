// main.go
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"path/filepath"

	"util/bezier"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Points []bezier.Point `json:"points"`
}

type RequestBody struct {
	Points    [][]float64 `json:"points"`
	Iteration int         `json:"iteration"`
}

func main() {
	fmt.Println("Started")
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))
	r.POST("/find-curve", handleCurve)
	r.Run(":8080")
}

func handleCurve(c *gin.Context) {
	var reqBody RequestBody

	if err := c.BindJSON(&reqBody); err != nil {
		fmt.Println("Error parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	curvePointer := bezier.BezierPoints{}

	for _, i := range reqBody.Points {
		curvePointer.InsertLast(bezier.Point{X: i[0], Y: i[1]})
	}

	directory := "dummy/"

	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		filePath := filepath.Join(directory, file.Name())
		err := os.Remove(filePath)
		if err != nil {
			fmt.Printf("Error removing %s: %v\n", filePath, err)
		}
	}

	curvePointer.InsertLast()
	curvePointer.DrawCurveBruteForce(int(math.Pow(2, float64(reqBody.Iteration)) + 1))
	poinnt := curvePointer.FindCurve(reqBody.Iteration)
	bezier.DrawSketch(poinnt, curvePointer, reqBody.Iteration)

	response := Response{
		Points: poinnt.List,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	fmt.Println(curvePointer)

	c.Data(http.StatusOK, "application/json", responseJSON)
}
