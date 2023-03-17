/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	openapi "github.com/syndis-software/iris-api/pkg/openapi"
)

// rootCmd represents the base command when called without any subcommands
var (
	// Used for flags
	// TODO: Description update
	rootCmd = &cobra.Command{
		Use:   "iris-api",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Printf("IN ROOT")
		// },
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

			api_key := viper.GetString("api_token")
			company := viper.GetString("company")
			host := viper.GetString("host")
			fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
			fmt.Printf("api_key: %v\n", api_key)
			fmt.Printf("company: %v\n", company)
			fmt.Printf("host: %v\n", host)

			ctx := cmd.Context()
			doer := ctx.Value("doer").(openapi.HttpRequestDoer)
			apiKeyIntercept, _ := openapi.NewSecurityProviderApiKey("x-api-key", api_key)
			client, _ := openapi.NewClient(host, openapi.WithRequestEditorFn(apiKeyIntercept.Intercept), openapi.WithHTTPClient(doer))
			ctx = context.WithValue(ctx, "client", client)
			cmd.SetContext(ctx)
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context, doer openapi.HttpRequestDoer) {
	ctx = context.WithValue(ctx, "doer", doer)
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().String("host", "https://app.vikin.gr", "Iris host")
	rootCmd.PersistentFlags().String("company", "", "Company value. Should look like Company-XXXX")
	viper.BindPFlag("company", rootCmd.PersistentFlags().Lookup("company"))
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.SetEnvPrefix("iris")
	viper.AutomaticEnv()

}
