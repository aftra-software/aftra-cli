/*
Copyright © 2023 Syndis ehf. <syndis@syndis.is>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	openapi "github.com/syndis-software/iris-api/pkg/openapi"
)

// configCmd represents the config command
var getTokenConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Get the config for a token",
	Long: `Get the config for a token.

The output is suitable for being piped into a file for future use`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		client := ctx.Value(clientKey).(*openapi.ClientWithResponses)
		tokenInfo, err := client.GetTokenInfoWithResponse(ctx)

		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error: %s", err)
			return
		}

		s, _ := strconv.Unquote(`"` + tokenInfo.JSON200.Config + `"`)
		fmt.Fprint(cmd.OutOrStdout(), s)
	},
}

func init() {
	getTokenCmd.AddCommand(getTokenConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
