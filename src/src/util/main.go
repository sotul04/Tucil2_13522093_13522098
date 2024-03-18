// main.go
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"util/bezier"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Points []bezier.Point `json:"points"`
	Time   float64        `json:"time"`
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
	r.POST("/devidenconquer", handleDnC)
	r.POST("/brute-force", handleBF)
	r.Run(":8080")
}

func handleDnC(c *gin.Context) {
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

	start := time.Now()

	poinnt := curvePointer.FindCurve(reqBody.Iteration)

	elapsedTime := time.Since(start)

	fmt.Println("Elapsed Time: ", elapsedTime)

	response := Response{
		Points: poinnt.List,
		Time:   float64(elapsedTime.Microseconds()),
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Data(http.StatusOK, "application/json", responseJSON)
}

func handleBF(c *gin.Context) {
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

	start := time.Now()

	poinnt := curvePointer.DrawCurveBruteForce(int(math.Pow(2, float64(reqBody.Iteration)) + 1))

	elapsedTime := time.Since(start)

	response := Response{
		Points: poinnt.List,
		Time:   float64(elapsedTime.Microseconds()),
	}

	responseJSON, err := json.Marshal(response)

	fmt.Println("Elapsed Time: ", elapsedTime)

	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Data(http.StatusOK, "application/json", responseJSON)
}
