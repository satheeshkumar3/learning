package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExternalAPIResponse struct {
	// Define the structure of the response from the external API
	// For example, if the external API returns a JSON object like:
	// {
	//     "message": "Hello, world!",
	//     "status": 200
	// }
	One string `json:"one"`
	Key string `json:"key"`
}

func main() {
	r := gin.Default()

	r.GET("/call-external-api", func(c *gin.Context) {
		externalAPIURL := "http://echo.jsontest.com/key/value/one/two" // Replace with the actual external API URL

		// Prepare the request to the external API
		req, err := http.NewRequest("GET", externalAPIURL, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		// Optionally, set headers if needed
		req.Header.Set("Accept", "application/json")

		// Call the external API
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call external API"})
			return
		}
		defer resp.Body.Close()

		// Check if the response status is OK
		if resp.StatusCode != http.StatusOK {
			c.JSON(resp.StatusCode, gin.H{"error": "Received non-OK response from external API"})
			return
		}

		// Decode the response from the external API
		var apiResponse ExternalAPIResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
			return
		}

		// Return the response from the external API to the client
		c.JSON(http.StatusOK, apiResponse)
	})

	// Start the server
	r.Run(":8080") // Default port is 8080
}
