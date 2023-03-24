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
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := ctx.Value(clientKey).(*openapi.ClientWithResponses)

		tokenInfo, err := openapi.DoGetTokenInfo(ctx, client)

		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", tokenInfo.Company)
		return nil

	},
}

func init() {
	getTokenCmd.AddCommand(getTokenCompanyCmd)
}
