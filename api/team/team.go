package team

import (
	"context"
	"fmt"
	"log"

	"github.com/cheynewallace/tabby"
	"github.com/google/uuid"
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

func CreateTeam(projectId string, teamName string, teamDescription string) {
	connection := api.GetConnection()

	ctx := context.Background()

	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	s := projectId
	parsedUuid, _ := uuid.Parse(s)

	teamToCreate := &core.WebApiTeam{
		Name:        &teamName,
		Description: &teamDescription,
		ProjectId:   &parsedUuid,
	}

	teamArgs := core.CreateTeamArgs{
		Team:      teamToCreate,
		ProjectId: &projectId,
	}

	responseValue, err := coreClient.CreateTeam(ctx, teamArgs)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Team ", *responseValue.Name, " created succesfully")
}

func DeleteTeam(projectId string, teamId string) {
	connection := api.GetConnection()

	ctx := context.Background()

	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	teamArgs := core.DeleteTeamArgs{
		ProjectId: &projectId,
		TeamId:    &teamId,
	}

	coreClient.DeleteTeam(ctx, teamArgs)
	fmt.Println("Project with Id: ", teamId, "was deleted succesfully!")
}
