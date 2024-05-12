package main

import (
	"context"
	"fmt"

	"worker/internal/application/executor"
	"worker/internal/domain/entities"
	"worker/internal/infrastructure/mongodb"
)

func main() {
	ctx := context.Background()

	cfg := mongodb.Config{
		IP:         "192.168.0.143",
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

	id := "6640c6293e6146acb9d6a9a6"
	run, err := app.JobUseCases.Run(ctx, id, map[string]entities.Variable{
		"token": {Type: "text", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0YWY0Yjg5My03NDlkLTQ4ZDAtODA1MS1iM2YzOTdhNWM0ZTAiLCJlbWFpbCI6Im10c3RydWV0ZWNoQGV4YW1wbGUuY29tIiwiZXhwIjoxNzE1NTA5Mzk3fQ.Nnt_3UVi7u8b4zXGW-dttX8mcpgArOp7t2fjdj08Pzg"},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(run)
}
