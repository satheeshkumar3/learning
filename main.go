package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExternalAPIResponse struct {
	One string `json:"one"`
	Key string `json:"key"`
}

func main() {
	r := gin.Default()

	r.GET("/call-external-api", func(c *gin.Context) {
		externalAPIURL := "http://echo.jsontest.com/key/value/one/two"

		req, err := http.NewRequest("GET", externalAPIURL, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call external API"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			c.JSON(resp.StatusCode, gin.H{"error": "Received non-OK response from external API"})
			return
		}

		var apiResponse ExternalAPIResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
			return
		}

		c.JSON(http.StatusOK, apiResponse)
	})

	r.Run(":8080")
}
