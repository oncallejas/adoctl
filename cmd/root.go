/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adoctl",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "Azure DevOps organization URL")
	//rootCmd.MarkFlagRequired("url")
	viper.BindPFlag("url", rootCmd.Flags().Lookup("url"))

	rootCmd.Flags().StringP("pat", "p", "", "Personal access token for authenticate with Azure DevOps Organization")
	//rootCmd.MarkFlagRequired("pat")
	viper.BindPFlag("pat", rootCmd.Flags().Lookup("pat"))
}
