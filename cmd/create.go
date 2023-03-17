/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	openapi "github.com/syndis-software/iris-api/pkg/openapi"
)

// createCmd represents the create command

var (
	uid        string
	name       string
	score      string
	detailsStr string

	createCmd = &cobra.Command{
		Use:   "create",
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

			company := viper.GetString("company")
			ctx := cmd.Context()
			client := ctx.Value("client").(*openapi.Client)
			resp, err := openapi.SendIt(ctx, client, company, opportunity)
			// TODO: handle errors here
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Error: %s", err)
				return
			}

			fmt.Printf("STATUS CODE %d", resp.StatusCode)
			switch code := resp.StatusCode; {
			case code == http.StatusUnauthorized:
				fmt.Fprintf(cmd.ErrOrStderr(), "Unauthorized")
			case code == http.StatusForbidden:
				fmt.Fprintf(cmd.ErrOrStderr(), "Forbidden")
			case code > 500:
				fmt.Fprintf(cmd.ErrOrStderr(), "Server Error")
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
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringVar(&uid, "uid", "", "Unique identifier for the opportunity")
	createCmd.Flags().StringVar(&name, "name", "", "Name of the opportunity")
	createCmd.Flags().StringVar(&score, "score", string(openapi.Unknown), "Risk score of the opportunity (critical, high, medium, low, info, none, unknown)")
	createCmd.Flags().StringVar(&detailsStr, "details", "", "Additional details. Comma separated key=value pairs.")
	createCmd.MarkPersistentFlagRequired("uid")
}
