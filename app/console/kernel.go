package console

import (
	"github.com/developbiao/webcore-demo/app/console/command/foo"
	"github.com/developbiao/webcore-demo/framework"
	"github.com/developbiao/webcore-demo/framework/cobra"
	"github.com/developbiao/webcore-demo/framework/command"
)

func RunCommand(container framework.Container) error {
	// Root command
	var rootCmd = &cobra.command{
		Use:   "web",
		Short: "web command",
		Long:  "web framework command tool",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	rootCmd.SetContainer(container)
	command.AddKernelCommands(rootCmd)
	AddAppCommand(rootCmd)
	return rootCmd.Execute()
}

func AddAppComand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(foo.FooCommand)
}
