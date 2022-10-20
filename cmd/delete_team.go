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
var delete_teamCmd = &cobra.Command{
	Use:   "team",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectId := viper.GetString("deleteTeamProjectId")
		teamId := viper.GetString("deleteTeamTeamId")
		team.DeleteTeam(projectId, teamId)
	},
}

func init() {
	deleteCmd.AddCommand(delete_teamCmd)

	delete_teamCmd.Flags().StringP("project-id", "p", "", "Project's Id where team to be deleted exists")
	delete_teamCmd.MarkFlagRequired("project-id")
	viper.BindPFlag("deleteTeamProjectId", delete_teamCmd.Flags().Lookup("project-id"))

	delete_teamCmd.Flags().StringP("team-id", "t", "", "Team's Id to be deleted")
	delete_teamCmd.MarkFlagRequired("team-id")
	viper.BindPFlag("deleteTeamTeamId", delete_teamCmd.Flags().Lookup("team-id"))
}
