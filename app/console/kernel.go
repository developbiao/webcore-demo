package console

import (
	"github.com/developbiao/webcore-demo/framework"
	"github.com/developbiao/webcore-demo/framework/cobra"

)

func RunCommand(container framework.Container ) error {
	// Root command
	var rootCmd = &cobra.command{
		Use: "web"

	}


}