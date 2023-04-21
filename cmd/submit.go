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

// logCmd represents the log command
var (
	submitCmd = &cobra.Command{
		Use:   "submit [scan-type] [scan-name]",
		Args:  cobra.MatchAll(cobra.ExactArgs(3), cobra.OnlyValidArgs),
		Short: "Submit a log message for the current token",
		Long: `Submit a log message for the current token
	
Log messages can be viewed against the token in the Aftra UI. This can provide a
useful feedback loop if you are configuring your token-using-application via
the token config in aftra. Simply pass in any string and it will appear there.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client := ctx.Value(clientKey).(*openapi.ClientWithResponses)
			scanType, scanName, message := args[0], args[1], args[2]

			switch {
			case ScanType(scanType) == syndis:

				// Log a single message
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
