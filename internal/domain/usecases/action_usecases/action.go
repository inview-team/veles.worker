package action_usecases

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"worker/internal/domain/entities"
)

type ActionUsecases struct {
	repo entities.ActionRepository
}

func New(repo entities.ActionRepository) (*ActionUsecases, error) {
	return &ActionUsecases{repo: repo}, nil
}

func (u *ActionUsecases) Register(actionType entities.ActionType, arguments map[string]entities.Variable) (string, error) {
	action, err := entities.NewJob(u.repo.NextID(), actionType, arguments)
	if err != nil {
		return "", fmt.Errorf("failed to register action: %v", err)
	}

	err = u.repo.Create(*action)
	if err != nil {
		return "", fmt.Errorf("failed to register action: %v", err)
	}
	return "", err
}

func (u *ActionUsecases) GetByID(actionId string) (entities.Action, error) {
	return u.repo.GetByID(actionId)
}

func (u *ActionUsecases) ExecuteHTTPAction(token string, action entities.Action) (*map[string]entities.Variable, bool, error) {
	client := &http.Client{}

	body := make(map[string]interface{})

	if len(action.Arguments) != 0 {
		for key, value := range action.Arguments {
			body[key] = value.Value
		}
	}
	b, err := json.Marshal(body)

	req, err := http.NewRequest(action.AdditionalParams["method"].(string), action.AdditionalParams["url"].(string), bytes.NewBuffer(b))
	if err != nil {
		return nil, false, fmt.Errorf("failed to execute http action: %v", err)
	}

	for key, value := range action.AdditionalParams["header"].(map[string]string) {
		req.Header.Set(key, value)
	}

	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, false, fmt.Errorf("failed to execute http action: %v", err)
	}
	defer resp.Body.Close()

	parsedResponse := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(parsedResponse)
	if err != nil {
		return nil, false, fmt.Errorf("failed to execute http action: %v", err)
	}

	output := make(map[string]entities.Variable)
	for key, value := range parsedResponse {
		output[key] = entities.Variable{Type: "string", Value: value}
	}
	return &output, true, nil
}

func (u *ActionUsecases) ExecuteCompareAction(action entities.Action) (*map[string]entities.Variable, bool, error) {
	firstValue, ok := action.Arguments["first_value"]
	if !ok {
		return nil, false, fmt.Errorf("failed to execute compare action: first value not provided")
	}
	secondValue, ok := action.Arguments["second_value"]
	if !ok {
		return nil, false, fmt.Errorf("failed to execute compare action: second value not provided")
	}

	operator, ok := action.Arguments["operator"]
	if !ok {
		return nil, false, fmt.Errorf("failed to execute compare action: operator not provided")
	}

	result, err := u.compare(firstValue.Value.(int), secondValue.Value.(int), entities.CompareOperations(operator.Value.(int)))
	if err != nil {
		return nil, false, fmt.Errorf("failed to execute compare action: %v", err)
	}

	output := make(map[string]entities.Variable)
	output["result"] = entities.Variable{Type: "bool", Value: result}
	return &output, true, nil
}

func (u *ActionUsecases) compare(firstValue, secondValue int, operator entities.CompareOperations) (bool, error) {
	switch operator {
	case entities.Equal:
		return firstValue == secondValue, nil
	case entities.More:
		return firstValue > secondValue, nil
	default:
		return false, fmt.Errorf("failed to compare values: unknown operation")
	}
}
