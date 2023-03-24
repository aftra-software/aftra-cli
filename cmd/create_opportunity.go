/*
Copyright © 2023 Syndis ehf. <syndis@syndis.is>

*/
package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	openapi "github.com/syndis-software/iris-api/pkg/openapi"
)

// opportunityCmd represents the opportunity command
var (
	uid        string
	name       string
	score      string
	detailsStr string

	opportunityCmd = &cobra.Command{
		Use:   "opportunity",
		Short: "Create internal opportunities inside Iris",
		Long: `Use the Iris API to create internal opportunities

	These will become part of the overall picture of your installation.
	You'll need an API key to make this happen`,
		Run: func(cmd *cobra.Command, args []string) {
			details := stringToMap(detailsStr)

			opportunity := openapi.CreateOpportunity{
				Name:    name,
				Uid:     uid,
				Score:   openapi.OpportunityScore(score),
				Details: details,
			}

			ctx := cmd.Context()
			client := ctx.Value(clientKey).(*openapi.ClientWithResponses)
			company := ctx.Value(companyKey).(string)
			resp, err := openapi.DoCreateOpportunity(ctx, client, company, opportunity)

			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Error: %s", err)
				return
			}

			switch code := resp.StatusCode; {
			case code == http.StatusUnauthorized:
				fmt.Fprintf(cmd.ErrOrStderr(), "Unauthorized")
			case code == http.StatusForbidden:
				fmt.Fprintf(cmd.ErrOrStderr(), "Forbidden")
			case code >= 500:
				fmt.Fprintf(cmd.ErrOrStderr(), "Server Error: %d", code)
			case code < 300:
				fmt.Fprintf(cmd.OutOrStdout(), "%s created", uid)
			default:
				fmt.Fprintf(cmd.OutOrStdout(), "Unrecognized status code %d", code)
			}

		},
	}
)

func stringToMap(str string) map[string]string {
	result := make(map[string]string)

	if str == "" {
		return result
	}
	// split the string into key-value pairs
	pairs := strings.Split(str, ",")
	// loop through each key-value pair
	for _, pair := range pairs {
		// split the pair into key and value
		kv := strings.Split(pair, "=")

		// skip empty key-value pairs
		if len(kv) != 2 || kv[0] == "" || kv[1] == "" {
			continue
		}

		// add the key-value pair to the map
		result[kv[0]] = kv[1]
	}

	return result
}

func init() {
	createCmd.AddCommand(opportunityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// opportunityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// opportunityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	opportunityCmd.Flags().StringVar(&uid, "uid", "", "Unique identifier for the opportunity")
	opportunityCmd.Flags().StringVar(&name, "name", "", "Name of the opportunity")
	opportunityCmd.Flags().StringVar(&score, "score", string(openapi.Unknown), "Risk score of the opportunity (critical, high, medium, low, info, none, unknown)")
	opportunityCmd.Flags().StringVar(&detailsStr, "details", "", "Additional details. Comma separated key=value pairs.")
	opportunityCmd.MarkFlagRequired("uid")
	opportunityCmd.MarkFlagRequired("name")
}
