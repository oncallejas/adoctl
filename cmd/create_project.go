/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var create_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called from create")
	},
}

func init() {
	createCmd.AddCommand(create_projectCmd)

	/* create_projectCmd.Flags().StringP("project-name", "pn", "", "Name of project to be created")
	create_projectCmd.MarkFlagRequired("name")
	viper.BindPFlag("project-name", create_projectCmd.Flags().Lookup("project-name")) */
}
