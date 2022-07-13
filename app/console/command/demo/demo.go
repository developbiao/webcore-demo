package demo

import (
	"fmt"

	"github.com/spf13/cobra"
)

// demo command
var DemoCommand = &cobra.Command{
	Use:   "demo",
	Short: "demo",
	RunE: func(c *cobra.Command, args []string) error {
		fmt.Println("hello demo command")
		return nil
	},
}
