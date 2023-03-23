/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	openapi "github.com/syndis-software/iris-api/pkg/openapi"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Submit a log message for the current token",
	Long: `Submit a log message for the current token
	
Log messages can be viewed against the token in the Iris UI. This can provide a
useful feedback loop if you are configuring your token-using-application via
the token config in iris. Simply pass in any string and it will appear there.
`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		client := ctx.Value(clientKey).(*openapi.ClientWithResponses)

		if len(args) > 0 {
			// Log a single message
			logs := make([]openapi.SubmitLogEvent, 0, 1)
			logs = append(logs, openapi.SubmitLogEvent{
				Message:   args[0],
				Timestamp: time.Now().UnixMilli(),
			})
			foo := openapi.SubmitLogsForTokenJSONRequestBody(logs)

			resp, err := client.SubmitLogsForToken(ctx, foo)

			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Error: %s", err)
				return
			}

			// TODO: Move this stuff into some generic error handler
			// TODO: Add tests around status codes of the command
			switch code := resp.StatusCode; {
			case code == http.StatusUnauthorized:
				fmt.Fprintf(cmd.ErrOrStderr(), "Unauthorized")
			case code == http.StatusForbidden:
				fmt.Fprintf(cmd.ErrOrStderr(), "Forbidden")
			case code >= 500:
				fmt.Fprintf(cmd.ErrOrStderr(), "Server Error: %d", code)
			case code <= 200:
			default:
				fmt.Fprintf(cmd.OutOrStdout(), "Unrecognized status code %d", code)
			}

		} else {
			// Assume we're in stdin mode.
		}

	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
