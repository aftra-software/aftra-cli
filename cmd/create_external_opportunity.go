/*
Copyright © 2023 Syndis ehf. <syndis@syndis.is>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	openapi "github.com/syndis-software/aftra-cli/pkg/openapi"
)

// opportunityExternalCmd represents the external opportunity command
var (
	externalUid   string
	externalName  string
	entityName    string
	externalScore int
	description   string

	opportunityExternalCmd = &cobra.Command{
		Use:          "external-opportunity",
		SilenceUsage: true,
		Short:        "Create external opportunities inside Aftra",
		Long: `Use the Aftra API to create external opportunities

	These will become part of the overall picture of your installation.
	You'll need an API key to make this happen`,
		RunE: func(cmd *cobra.Command, args []string) error {

			opportunity := openapi.CreateExternalOpportunity{
				Name:        externalName,
				EntityName:  entityName,
				Uid:         externalUid,
				Score:       openapi.OpportunityScore(externalScore),
				Description: description,
			}

			ctx := cmd.Context()
			client := ctx.Value(clientKey).(*openapi.ClientWithResponses)
			company := ctx.Value(companyKey).(string)
			err := openapi.DoCreateExternalOpportunity(ctx, client, company, opportunity)

			if err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "%s created\n", externalUid)
			return err
		},
	}
)

func init() {
	createCmd.AddCommand(opportunityExternalCmd)
	opportunityExternalCmd.Flags().StringVar(&externalUid, "uid", "", "Unique identifier for the opportunity")
	opportunityExternalCmd.Flags().StringVar(&externalName, "name", "", "Name of the opportunity")
	opportunityExternalCmd.Flags().StringVar(&entityName, "entity", "", "Name of the entity that will be linked to this opportunity")
	opportunityExternalCmd.Flags().IntVar(&externalScore, "score", -1, "Risk score of the opportunity (critical (5), high (4), medium (3), low (2), info (1), none (0), unknown (-1))")
	opportunityExternalCmd.Flags().StringVar(&description, "description", "", "The description of the opportunity")
	opportunityExternalCmd.MarkFlagRequired("uid")
	opportunityExternalCmd.MarkFlagRequired("name")
	opportunityExternalCmd.MarkFlagRequired("entity")
}