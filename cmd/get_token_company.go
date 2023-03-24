/*
Copyright © 2023 Syndis ehf. <syndis@syndis.is>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	openapi "github.com/syndis-software/iris-api/pkg/openapi"
)

// companyCmd represents the company command
var getTokenCompanyCmd = &cobra.Command{
	Use:   "company",
	Short: "Get the company associated with this token",
	Long: `Get the company associated with this token
	
Some commands require a company id. Use this command to get the company from the 
API. This can be then set as the environment variable IRIS_COMPANY for future use.
`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		client := ctx.Value(clientKey).(*openapi.ClientWithResponses)
		tokenInfo, err := client.GetTokenInfoWithResponse(ctx)

		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error: %s", err)
			return
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", tokenInfo.JSON200.Company)
	},
}

func init() {
	getTokenCmd.AddCommand(getTokenCompanyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// companyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// companyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
