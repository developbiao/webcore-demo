package middleware

import (
	"log"
	"time"

	"github.com/developbiao/webcore-demo/framework/gin"
)

func Cost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record the start time
		start := time.Now()

		// Use next execute logic bussiness
		c.Next()

		// Record the end time
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %s, cost: %d", c.Request.RequestURI, cost.Milliseconds())
	}
}
