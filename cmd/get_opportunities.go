// Get opportunities usage

// aftra-cli list opportunities --updated-since=DT --limit=100

// aftra-cli update-resolution $UID $RESOLUTION

/*
Copyright © 2023 Syndis ehf. <syndis@syndis.is>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	openapi "github.com/syndis-software/aftra-cli/pkg/openapi"
)

// getCmd represents the get command

func GetTime(s string) (time.Time, error) {
	// Taken from proposed implementation of parsing times in cobra
	// https://github.com/spf13/pflag/pull/348/files
	s = strings.TrimSpace(s)
	formats := []string{time.RFC3339Nano, time.RFC1123Z}
	for _, f := range formats {
		v, err := time.Parse(f, s)
		if err != nil {
			continue
		}
		return v, nil
	}

	formatsString := ""
	for i, f := range formats {
		if i > 0 {
			formatsString += ", "
		}
		formatsString += fmt.Sprintf("`%s`", f)
	}

	return time.Time{}, fmt.Errorf("invalid time format `%s` must be one of: %s", s, formatsString)
}

var (
	limit        int
	updatedSince string

	getOpportunitiesCmd = &cobra.Command{
		Use:   "opportunities ",
		Short: "Get opportunities",
		Long: `Get filtered opportunities.

Output is JSON format`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client := ctx.Value(clientKey).(*openapi.ClientWithResponses)
			company := ctx.Value(companyKey).(string)

			lastUpdatedGte, err := GetTime(updatedSince)

			if err != nil {
				return err
			}

			opportunities, err := openapi.DoGetOpportunities(ctx, client, company, &lastUpdatedGte, &limit)

			if err != nil {
				return err
			}

			for _, oppo := range opportunities.Opportunities {
				txt, err := json.Marshal(oppo)
				if err != nil {
					return err
				}
				fmt.Fprintf(cmd.OutOrStdout(), "%s\n", txt)

			}
			return nil
		},
	}
)

func init() {
	getCmd.AddCommand(getOpportunitiesCmd)
	getOpportunitiesCmd.Flags().IntVar(&limit, "limit", 100, "Max number of opportunities to retrieve")
	getOpportunitiesCmd.Flags().StringVar(&updatedSince, "updated-since", "2020-01-01T00:00:00Z", "Only fetch opportunities updated since")

}
