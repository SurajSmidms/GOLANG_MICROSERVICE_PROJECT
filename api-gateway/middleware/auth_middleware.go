package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// validateResponse represents the expected JSON response shape from Auth Service /validate
type validateResponse map[string]interface{}

// AuthMiddleware verifies the incoming Authorization header by calling Auth Service /validate.
// On success it adds X-User (username/email) header and continues.
// Public endpoints should skip this middleware (we will not attach it to /auth/* public routes).
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		// Build request to auth service validate endpoint
		authServiceURL := os.Getenv("AUTH_SERVICE_URL")
		if authServiceURL == "" {
			log.Println("AUTH_SERVICE_URL not set")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gateway misconfigured"})
			c.Abort()
			return
		}

		validateURL := strings.TrimRight(authServiceURL, "/") + "/validate"

		req, err := http.NewRequest("GET", validateURL, nil)
		if err != nil {
			log.Println("failed to create request to auth validate:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			c.Abort()
			return
		}
		// Forward the Authorization header
		req.Header.Set("Authorization", authHeader)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("auth service unreachable:", err)
			c.JSON(http.StatusBadGateway, gin.H{"error": "auth service unavailable"})
			c.Abort()
			return
		}
		defer resp.Body.Close()

		bodyBytes, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			// forward auth error
			var errMsg interface{}
			_ = json.Unmarshal(bodyBytes, &errMsg)
			c.JSON(resp.StatusCode, gin.H{"error": errMsg})
			c.Abort()
			return
		}

		// parse response to extract user identity
		var vr validateResponse
		if err := json.Unmarshal(bodyBytes, &vr); err != nil {
			// If parsing fails, still allow but without user header
			log.Println("warning: failed to parse auth validate response:", err)
			c.Next()
			return
		}

		// try to extract common keys (user, username, email)
		var userVal string
		if v, ok := vr["user"].(string); ok && v != "" {
			userVal = v
		} else if v, ok := vr["username"].(string); ok && v != "" {
			userVal = v
		} else if v, ok := vr["email"].(string); ok && v != "" {
			userVal = v
		} else if msg, ok := vr["message"].(string); ok && strings.Contains(strings.ToLower(msg), "token valid") {
			// fallback: if token valid and email provided in response? ignore
			userVal = ""
		}

		// attach identity header for downstream services if we found one
		if userVal != "" {
			c.Request.Header.Set("X-User", userVal)
			c.Set("X-User", userVal)
		}

		// if auth is OK, continue
		c.Next()
	}
}
