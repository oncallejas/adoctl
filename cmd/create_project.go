/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/oncallejas/adoctl/api/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		projectName := viper.GetString("project-name")
		projectDescription := viper.GetString("project-description")
		projectSourceControlType := viper.GetString("project-source-control-type")
		projectProcessId := viper.GetString("project-process-id")

		project.CreateProject(&projectName, &projectDescription, projectSourceControlType, projectProcessId)
	},
}

func init() {
	createCmd.AddCommand(create_projectCmd)

	// Flag for project's name
	create_projectCmd.Flags().StringP("name", "n", "", "Name of the project to be created")
	create_projectCmd.MarkFlagRequired("name")
	viper.BindPFlag("project-name", create_projectCmd.Flags().Lookup("name"))

	// Flag for project's description
	create_projectCmd.Flags().StringP("description", "d", "", "Projects description to be created")
	create_projectCmd.MarkFlagRequired("description")
	viper.BindPFlag("project-description", create_projectCmd.Flags().Lookup("description"))

	// Flag to set project's visibility
	// create_projectCmd.Flags().StringP("visibility", "v", "private", "Visibility of the project")
	// create_projectCmd.MarkFlagRequired("visibility")
	// viper.BindPFlag("project-visibility", create_projectCmd.Flags().Lookup("visibility"))

	// Flag to set the version control type of the project
	create_projectCmd.Flags().StringP("source-control-type", "s", "", "Source control type of the project")
	create_projectCmd.MarkFlagRequired("source-control-type")
	viper.BindPFlag("project-source-control-type", create_projectCmd.Flags().Lookup("source-control-type"))

	// Flag for process id to be relationated with project
	create_projectCmd.Flags().StringP("process-id", "i", "", "Process Id to be used to create project")
	create_projectCmd.MarkFlagRequired("process-id")
	viper.BindPFlag("project-process-id", create_projectCmd.Flags().Lookup("process-id"))
}
