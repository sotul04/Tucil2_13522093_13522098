// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Points []Point `json:"points"`
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

	// BindJSON digunakan untuk mengurai body permintaan JSON ke struct RequestBody
	if err := c.BindJSON(&reqBody); err != nil {
		fmt.Println("Error parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	curvePointer := BezierPoints{}

	for _, i := range reqBody.Points {
		curvePointer.insertLast(Point{i[0], i[1]})
	}

	curvePointer.insertLast()
	curvePointer.drawCurveBruteForce()
	poinnt := curvePointer.findCurve(reqBody.Iteration)
	drawSketch(poinnt, curvePointer, reqBody.Iteration)

	response := Response{
		Points: poinnt.list,
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
