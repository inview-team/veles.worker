package worker

import (
	"worker/internal/domain/entities"
)

func main() {

	balanceParams := make(map[string]interface{})
	balanceParams["url"] = "http://192.168.0.143:30002/api/v1/wallet"
	balanceParams["method"] = "GET"

	balanceArgs := map[string]entities.Variable{
		"token": {Type: "text", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0YWY0Yjg5My03NDlkLTQ4ZDAtODA1MS1iM2YzOTdhNWM0ZTAiLCJlbWFpbCI6Im10c3RydWV0ZWNoQGV4YW1wbGUuY29tIiwiZXhwIjoxNzE1NTA5Mzk3fQ.Nnt_3UVi7u8b4zXGW-dttX8mcpgArOp7t2fjdj08Pzg"},
	}
	balanceAction, _ := entities.NewAction("1", entities.Request, balanceArgs, balanceParams)

	validateArgs := map[string]entities.Variable{
		"balance":  {Type: "number", Value: nil},
		"amount":   {Type: "number", Value: nil},
		"operator": {Type: "enum", Value: entities.More},
	}
	validateAction, _ := entities.NewAction("2", entities.Validate, validateArgs, nil)

	seq := map[entities.ActionID]entities.ActionID{
		balanceAction.Id: validateAction.Id,
	}

	output := entities.JobOutput{
		Ask: entities.Output{
			Message:  "Get amount for transfer",
			Variable: entities.Variable{},
		},
		OnSuccess: entities.Output{
			Message:  "You can make transfer",
			Variable: entities.Variable{},
		},
		OnFailure: entities.Output{
			Message:  "Not enough money",
			Variable: entities.Variable{},
		},
	}
	job, err := entities.NewJob("1", balanceAction.Id, seq, output)
}
