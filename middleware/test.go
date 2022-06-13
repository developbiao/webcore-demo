package middleware

import (
	"fmt"

	"github.com/developbiao/webcore-demo/framework"
)

func Test1() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test01")
		c.Next()
		fmt.Println("middleware post test01")
		return nil
	}
}

func Test2() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test02")
		c.Next()
		fmt.Println("middleware post test02")
		return nil
	}
}

func Test3() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test03")
		c.Next()
		fmt.Println("middleware post test03")
		return nil
	}
}
