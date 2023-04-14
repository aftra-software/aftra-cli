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
var getScanConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Get the config for a scanner",
	Long: `Get the config for a scanner.

The output is suitable for being piped into a file for future use`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := ctx.Value(clientKey).(*openapi.ClientWithResponses)

		tokenInfo, err := openapi.DoGetTokenInfo(ctx, client)

		if err != nil {
			return err
		}

		s, _ := strconv.Unquote(`"` + tokenInfo.Config + `"`)
		fmt.Fprint(cmd.OutOrStdout(), s)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getScanConfigCmd)
}
