/*
Copyright © 2023 Syndis ehf. <syndis@syndis.is>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	openapi "github.com/syndis-software/iris-api/pkg/openapi"
)

// getCmd represents the get command

var getConfigCmd = &cobra.Command{
	Use:   "config [scan-type] [scan-name]",
	Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Short: "Get the config for a scanner",
	Long: `Get the config for a scanner.

The output is suitable for being piped into a file for future use`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := ctx.Value(clientKey).(*openapi.ClientWithResponses)
		scanType, scanName := args[0], args[1]

		switch {
		case ScanType(scanType) == syndis:
			configInfo, err := openapi.DoGetScanConfig(ctx, scanName, client)
			if err != nil {
				return err
			}
			fmt.Fprint(cmd.OutOrStdout(), configInfo)
			return nil
		default:
			return fmt.Errorf("unrecognised scan type %s", scanType)
		}

	},
}

func init() {
	getCmd.AddCommand(getConfigCmd)
}
