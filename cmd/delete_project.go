/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/google/uuid"
	"github.com/oncallejas/adoctl/api/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// projectCmd represents the project command
var delete_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := viper.GetString("projectId")
		parsedUuid, _ := uuid.Parse(s)

		project.DeleteProject(&parsedUuid)
	},
}

func init() {
	deleteCmd.AddCommand(delete_projectCmd)

	delete_projectCmd.Flags().StringP("project-id", "i", "", "Project Id to be deleted")
	delete_projectCmd.MarkFlagRequired("project-id")
	viper.BindPFlag("projectId", delete_projectCmd.Flags().Lookup("project-id"))
}
