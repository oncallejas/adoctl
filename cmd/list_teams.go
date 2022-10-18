/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/oncallejas/adoctl/api/team"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// teamsCmd represents the teams command
var list_teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		teamProjectId := viper.GetString("project-id")
		team.ListTeams(&teamProjectId)
	},
}

func init() {
	listCmd.AddCommand(list_teamsCmd)

	list_teamsCmd.Flags().StringP("project-id", "i", "", "project id to query teams")
	list_teamsCmd.MarkFlagRequired("project-id")
	viper.BindPFlag("project-id", list_teamsCmd.Flags().Lookup("project-id"))
}
