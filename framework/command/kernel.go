package command

import "github.com/developbiao/webcore-demo/framework/cobra"

func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(DemoCommand)
}
