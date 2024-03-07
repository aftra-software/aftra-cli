// Get opportunities usage

// aftra-cli list opportunities --updated-since=DT --limit=100

// aftra-cli update-resolution $UID $RESOLUTION

/*
Copyright © 2023 Syndis ehf. <syndis@syndis.is>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/syndis-software/aftra-cli/pkg/openapi"
	"golang.org/x/exp/slices"
)

// getCmd represents the get command

func validateResolution(resolution string) error {

	possible := []openapi.OpportunityResolution{
		openapi.AcceptedRisk,
		openapi.Unacknowledged,
		openapi.FalsePositive,
		openapi.Resolved,
	}
	if !slices.Contains(possible, openapi.OpportunityResolution(resolution)) {
		return fmt.Errorf("resolution must be one of %v. Got %s", possible, resolution)
	}
	return nil

}

var (
	commentStr string

	updateResolutionsCmd = &cobra.Command{
		Use:   "resolution [uid] [status]",
		Short: "Update resolution",
		Long:  `Update a resolution of an opportunity.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			uidStr, resolutionStr := args[0], args[1]

			err := validateResolution(resolutionStr)
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			client := ctx.Value(clientKey).(*openapi.ClientWithResponses)
			company := ctx.Value(companyKey).(string)

			return openapi.DoPostResolution(
				ctx,
				client,
				company,
				uidStr,
				openapi.ResolutionUpdate{
					Comment:    &commentStr,
					Resolution: openapi.OpportunityResolution(resolutionStr),
				},
			)

		},
	}
)

func init() {
	updateCmd.AddCommand(updateResolutionsCmd)
	updateResolutionsCmd.Flags().StringVar(&commentStr, "comment", "", "Comment associated with the resolution")

}
