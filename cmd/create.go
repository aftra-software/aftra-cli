/*
Copyright © 2023 Syndis ehf. <syndis@syndis.is>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command

var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create entities inside Iris",
		Long: `Use the Iris API to create things.

You'll need an API key to make this happen`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.ErrOrStderr(), "Error: must also specify a resource. eg opportunity")
		},
	}
)

func init() {
	rootCmd.AddCommand(createCmd)
}
