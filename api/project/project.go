package project

import (
	"context"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/oncallejas/greetctl/api"
)

func ListProjects() {
	config, err := api.LoadConfig("$HOME")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	organizationUrl := config.ADO_URL
	personalAccessToken := config.ADO_TOKEN

	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)

	ctx := context.Background()

	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	responseValue, err := coreClient.GetProjects(ctx, core.GetProjectsArgs{})
	if err != nil {
		log.Fatal(err)
	}

	index := 0
	for responseValue != nil {
		for _, teamProjectReference := range (*responseValue).Value {
			log.Printf("Name[%v] = %v", index, *teamProjectReference.Name)
			index++
		}

		if responseValue.ContinuationToken != "" {
			// Get next page of team projects
			projectArgs := core.GetProjectsArgs{
				ContinuationToken: &responseValue.ContinuationToken,
			}
			responseValue, err = coreClient.GetProjects(ctx, projectArgs)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			responseValue = nil
		}
	}
}
