/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	openapi "github.com/syndis-software/iris-api/pkg/openapi"
)

// tokenCmd represents the token command
var getTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Get token information for currently active token",
	Long: `Get token information for currently active token

Without args, this will output the full json text representing the token.

Supply "config" or "company" arguments to get escaped versions for 
use in setup scripts.
`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		client := ctx.Value(clientKey).(*openapi.ClientWithResponses)
		tokenInfo, err := openapi.DoFetchToken(ctx, client)

		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error: %s", err)
			return
		}

		s, err := json.MarshalIndent(tokenInfo.JSON200, "", "\t")

		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Failed to unmarshall: %s", err)
			return
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", string(s))
	},
}

func init() {
	getCmd.AddCommand(getTokenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tokenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tokenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
