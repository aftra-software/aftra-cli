/*
Copyright © 2023 Syndis ehf. <syndis@syndis.is>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	openapi "github.com/syndis-software/aftra-cli/pkg/openapi"
)

var (
	submitCmd_filename string
	submitCmd_message  string

	submitCmd = &cobra.Command{
		Use:   "submit [scan-type] [scan-name] [scan-result]",
		Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
		Short: "Submit a json formatted scan result",
		Long: `Submit a json formatted scan result
	
Submit a scan result in the format for given scan-type. For example in
nessus format for syndis scans.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client := ctx.Value(clientKey).(*openapi.ClientWithResponses)
			scanType, scanName := args[0], args[1]

			switch {
			case ScanType(scanType) == syndis:

				var scans []openapi.SyndisInternalScanEvent
				// Submit a set of scan events
				if submitCmd_filename != "" {
					jsonFile, err := os.Open(submitCmd_filename)
					// if we os.Open returns an error then handle it
					if err != nil {
						return err
					}
					// defer the closing of our jsonFile so that we can parse it later on
					defer jsonFile.Close()
					byteValue, _ := ioutil.ReadAll(jsonFile)
					err = json.Unmarshal(byteValue, &scans)
					if err != nil {
						return err
					}
				} else {
					err := json.Unmarshal([]byte(submitCmd_message), &scans)
					if err != nil {
						return err
					}
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

	submitCmd.Flags().StringVarP(&submitCmd_filename, "filename", "f", "", "JSON file to submit")
	submitCmd.Flags().StringVarP(&submitCmd_message, "message", "m", "", "JSON string to submit")

	// Want this, but no way to clear flag values between tests at the moment
	// https://github.com/spf13/cobra/issues/1180
	// submitCmd.MarkFlagsMutuallyExclusive("filename", "message")

	rootCmd.AddCommand(submitCmd)

}
