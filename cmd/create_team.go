/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/oncallejas/adoctl/api/team"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// teamCmd represents the team command
var create_teamCmd = &cobra.Command{
	Use:   "team",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectId := viper.GetString("createTeamProjectId")
		teamName := viper.GetString("teamName")
		teamDescription := viper.GetString("teamDescription")

		team.CreateTeam(projectId, teamName, teamDescription)
	},
}

func init() {
	createCmd.AddCommand(create_teamCmd)

	// Flag to set team name
	create_teamCmd.Flags().StringP("name", "n", "", "Team's name")
	create_teamCmd.MarkFlagRequired("name")
	viper.BindPFlag("teamName", create_teamCmd.Flags().Lookup("name"))

	// Flag to set team description
	create_teamCmd.Flags().StringP("description", "d", "", "Team's description")
	create_teamCmd.MarkFlagRequired("description")
	viper.BindPFlag("teamDescription", create_teamCmd.Flags().Lookup("description"))

	// Flag to set project where team will be created
	create_teamCmd.Flags().StringP("project-id", "i", "", "Project Id where team will be created")
	create_teamCmd.MarkFlagRequired("project-id")
	viper.BindPFlag("createTeamProjectId", create_teamCmd.Flags().Lookup("project-id"))
}
