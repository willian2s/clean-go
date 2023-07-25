package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	ips          map[string]int
	banDurations map[string]time.Time
)

func RateLimit(c *gin.Context) {
	ip := c.ClientIP()
	now := time.Now()

	if ips == nil {
		ips = make(map[string]int)
	}
	if banDurations == nil {
		banDurations = make(map[string]time.Time)
	}

	count, ok := ips[ip]
	if !ok {
		count = 0
	}

	if count >= 10 {
		banExpiration, banExists := banDurations[ip]
		if banExists && banExpiration.After(now) {
			c.Abort()
			c.String(http.StatusServiceUnavailable, "You are currently blocked. Please try again later.")
			return
		}

		delete(ips, ip)
		delete(banDurations, ip)

		banDuration := time.Duration(1) * time.Minute // Define a duração do bloqueio, por exemplo, 1 minuto
		banDurations[ip] = now.Add(banDuration)
		fmt.Println("IP blocked")
		c.Abort()
		c.String(http.StatusServiceUnavailable, "You have been automatically blocked")
		return
	}

	ips[ip] = count + 1
}
