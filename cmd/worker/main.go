package main

import (
	"context"
	"fmt"

	"github.com/inview-team/veles.worker/pkg/application/executor"
	"github.com/inview-team/veles.worker/pkg/domain/entities"
	"github.com/inview-team/veles.worker/pkg/infrastructure/mongodb"
)

func main() {
	ctx := context.Background()

	cfg := mongodb.Config{
		IP:         "127.0.0.1",
		Port:       27017,
		User:       "root",
		Password:   "password",
		AuthSource: "admin",
	}

	app, err := executor.NewApp(ctx, cfg)
	if err != nil {
		fmt.Println(err)
	}
	/*
		arguments := map[string]entities.Variable{
			"token":  {Type: "text", Value: nil},
			"to":     {Type: "text", Value: nil},
			"amount": {Type: "number", Value: nil},
		}
		output := []string{
			"balance",
		}

		params := make(map[string]interface{})
		params["url"] = "http://192.168.0.143:30002/api/v1/transfer"
		params["method"] = "POST"
	*/

	// id, err := app.ActionUseCases.Register(ctx, "Get Balance", entities.Request, arguments, output, params)
	// action, err := app.ActionUseCases.GetByID(ctx, "6640ae3412f4c109e563dd60")
	if err != nil {
		fmt.Println(err)
	}

	/*
		actions := []entities.ActionInformation{
			{Id: "6640ae3412f4c109e563dd60", Output: []string{"balance"}},
			{Id: "6640ae3412f4c109e563dd60", Output: []string{"balance"}},
			{Id: "6640ae3412f4c109e563dd60", Output: []string{"balance"}},
		}

		output := entities.JobOutput{
			Ask:       entities.Output{Message: "Not enough arguments", Type: entities.Ask},
			OnSuccess: entities.Output{Message: "Successful get Balance", Type: entities.Success},
			OnFailure: entities.Output{Message: "Failed to get Balance", Type: entities.Failure},
		}

		id, err := app.JobUseCases.Create(ctx, "Get Balance", actions, output)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)

	*/

	id := "6640d74b84049cbd5e07a70c"
	run, err := app.JobUseCases.Run(ctx, id, map[string]entities.Variable{
		"token":          {Type: "text", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0YWY0Yjg5My03NDlkLTQ4ZDAtODA1MS1iM2YzOTdhNWM0ZTAiLCJlbWFpbCI6Im10c3RydWV0ZWNoQGV4YW1wbGUuY29tIiwiZXhwIjoxNzE1NTMwNzQ1fQ.I9JaQIET0Yu41TnzhfvDl5eXwj9DfDqZW8_LHNKj3ZQ"},
		"from_wallet_id": {Type: "text", Value: "6902fb11-9e7a-4a0d-aa45-de428ba48861"},
		"to_wallet_id":   {Type: "text", Value: "214c4a21-52cb-409e-99cf-285c8b0fb676"},
		"amount":         {Type: "number", Value: 100},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(run)

	/*
		actions := []entities.ActionInformation{
			{Id: "6640b20ea992c6f3eb368046", Output: []string{"id", "receiver_id"}},
		}

		output := entities.JobOutput{
			Ask:       entities.Output{Message: "Provide arguments", Type: entities.Ask},
			OnSuccess: entities.Output{Message: "Successful transfer money", Type: entities.Success},
			OnFailure: entities.Output{Message: "Failed transfer money", Type: entities.Failure},
		}

		id, err := app.JobUseCases.Create(ctx, "Transfer Money", actions, output)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)
	*/
}
