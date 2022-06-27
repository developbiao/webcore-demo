package middleware

import "github.com/developbiao/webcore-demo/framework"

// Recovery caputures panics exception functions
func Recovery() framework.ControllerHandler {
	return func(c *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				c.SetStatus(500).Json(err)
			}
		}()
		// Using next execute logic bussiness
		c.Next()
		return nil
	}
}
