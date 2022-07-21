package foo

import (
	"fmt"

	"github.com/developbiao/webcore-demo/framework/cobra"
)

var FooCommand = &cobra.Command{
	Use:   "foo",
	Short: "foo",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.Container()
		fmt.Println(container)
		return nil
	},
}
