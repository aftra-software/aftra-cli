/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

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
			fmt.Println("create called")
			fmt.Printf("%v\n", uid)
			fmt.Printf("%v\n", host)

			details := stringToMap(detailsStr)

			opportunity := openapi.CreateOpportunity{
				Name:    name,
				Uid:     uid,
				Score:   openapi.OpportunityScore(score),
				Details: details,
			}

			ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*300))

			api_key = viper.GetString("api_token")
			company = viper.GetString("company")
			fmt.Printf("api_key: %s\n", api_key)
			fmt.Printf("company: %s\n", company)
			resp, err := openapi.SendIt(ctx, host, api_key, company, opportunity)
			fmt.Printf("err: %v", err)
			fmt.Printf("response: %v", resp)
			fmt.Printf("status: %v", resp.StatusCode)
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
}
