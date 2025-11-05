package utils

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// ProxyRequest forwards the incoming request to targetURL (full URL) and writes back the response.
func ProxyRequest(c *gin.Context, targetURL string) {
	// Build request with query string preserved
	u, err := url.Parse(targetURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid target URL"})
		return
	}
	u.RawQuery = c.Request.URL.RawQuery

	// Create new request for backend
	req, err := http.NewRequest(c.Request.Method, u.String(), c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create proxy request"})
		return
	}

	// Copy headers from original request
	for k, v := range c.Request.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}

	// Use a client to perform request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("❌ Proxy error to %s: %v", u.String(), err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "service unavailable"})
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for k, v := range resp.Header {
		for _, vv := range v {
			c.Writer.Header().Add(k, vv)
		}
	}

	c.Status(resp.StatusCode)
	// Stream body
	_, _ = io.Copy(c.Writer, resp.Body)
}
