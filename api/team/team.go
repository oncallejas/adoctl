package team

import (
	"context"
	"log"

	"github.com/cheynewallace/tabby"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/oncallejas/adoctl/api"
)

func ListTeams(teamProjectId *string) {
	connection := api.GetConnection()

	ctx := context.Background()

	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	teamArgs := core.GetTeamsArgs{}
	teamArgs.ProjectId = teamProjectId
	responseValue, err := coreClient.GetTeams(ctx, teamArgs)
	if err != nil {
		log.Fatal(err)
	}

	index := 0
	t := tabby.New()
	t.AddHeader("Id", "Name")
	for _, teamReference := range *responseValue {
		t.AddLine(*teamReference.Id, *teamReference.Name)
		index++
	}
	t.Print()
}
