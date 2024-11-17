package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TimeSeriesData represents a single entry in a time series (e.g., temperature readings).
type TimeSeriesData struct {
	Timestamp time.Time `json:"timestamp"` // The time of the measurement
	Value     float64   `json:"value"`     // The value of the metric (e.g., temperature)
}

// TimeSeries represents a collection of time-series data points.
type TimeSeries struct {
	Name   string           `json:"name"`   // The name or type of the time series (e.g., "temperature")
	Points []TimeSeriesData `json:"points"` // A slice of time-series data points
}

func main() {
	// Initialize the time series data inside the main function
	timeSeries := TimeSeries{
		Name: "Temperature",
		Points: []TimeSeriesData{
			{
				Timestamp: time.Now().Add(-5 * time.Minute),
				Value:     22.5,
			},
			{
				Timestamp: time.Now().Add(-4 * time.Minute),
				Value:     22.7,
			},
			{
				Timestamp: time.Now().Add(-3 * time.Minute),
				Value:     22.8,
			},
			{
				Timestamp: time.Now().Add(-2 * time.Minute),
				Value:     23.0,
			},
			{
				Timestamp: time.Now().Add(-1 * time.Minute),
				Value:     23.1,
			},
		},
	}

	// Create a new Gin router instance
	router := gin.Default()

	// Define the route and handler
	router.GET("/temp", func(c *gin.Context) {
		// Respond with the timeSeries object as JSON
		c.IndentedJSON(http.StatusOK, timeSeries)
	})

	// Start the Gin server on localhost:8080
	router.Run("localhost:8080")
}
