/*
Copyright © 2023 Syndis ehf. <syndis@syndis.is>

*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	openapi "github.com/syndis-software/aftra-api/pkg/openapi"
)

var (
	submitCmd = &cobra.Command{
		Use:   "submit [scan-type] [scan-name] [scan-result]",
		Args:  cobra.MatchAll(cobra.ExactArgs(3), cobra.OnlyValidArgs),
		Short: "Submit a json formatted scan result",
		Long: `Submit a json formatted scan result
	
Submit a scan result in the format for given scan-type. For example in
nessus format for syndis scans.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client := ctx.Value(clientKey).(*openapi.ClientWithResponses)
			scanType, scanName, message := args[0], args[1], args[2]

			switch {
			case ScanType(scanType) == syndis:

				// Submit a set of scan events

				var scans []openapi.SyndisInternalScanEvent
				err := json.Unmarshal([]byte(message), &scans)

				if err != nil {
					return err
				}

				resp, err := client.SubmitScanResults(ctx, scanName, scans)

				if err != nil {
					return err
				}

				return openapi.CheckStatus(resp)
			default:
				return fmt.Errorf("unrecognised scan type %s", scanType)
			}

		},
	}
)

func init() {
	rootCmd.AddCommand(submitCmd)

}
