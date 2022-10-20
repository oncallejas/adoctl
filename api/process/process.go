package process

import (
	"context"
	"log"

	"github.com/cheynewallace/tabby"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/oncallejas/adoctl/api"
)

func ListAllProcesses() {
	connection := api.GetConnection()

	ctx := context.Background()

	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	processesArgs := core.GetProcessesArgs{}

	responseValue, err := coreClient.GetProcesses(ctx, processesArgs)
	if err != nil {
		log.Fatalln(err)
	}

	t := tabby.New()
	t.AddHeader("Id", "Name", "Process Type", "Is Default")
	for _, processReference := range *responseValue {
		t.AddLine(*processReference.Id, *processReference.Name, *processReference.Type, *processReference.IsDefault)
	}
	t.Print()
}
