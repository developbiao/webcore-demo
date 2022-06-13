package middleware

import (
	"log"
	"time"

	"github.com/developbiao/webcore-demo/framework"
)

func Cost() framework.ControllerHandler {
	return func(c *framework.Context) error {
		// Record the start time
		start := time.Now()

		// Use next execute logic bussiness
		c.Next()

		// Record the end time
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %s, cost: %d", c.GetRequest().RequestURI, cost.Milliseconds())
		return nil
	}
}
