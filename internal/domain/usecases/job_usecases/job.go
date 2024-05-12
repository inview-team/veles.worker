package job_usecases

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"worker/internal/domain/entities"
)

type JobUsecases struct {
	repo  entities.JobRepository
	aRepo entities.ActionRepository
}

func New(repo entities.JobRepository, aRepo entities.ActionRepository) JobUsecases {
	return JobUsecases{
		repo:  repo,
		aRepo: aRepo,
	}
}

func (u *JobUsecases) Create(actions []entities.ActionInformation, output entities.JobOutput) (string, error) {
	job, err := entities.NewJob(u.repo.NextID(), actions, output)
	if err != nil {
		return "", fmt.Errorf("failed to create job: %v", err)
	}

	err = u.repo.Create(*job)
	if err != nil {
		return "", fmt.Errorf("failed to create job: %v", err)
	}
	return string(job.Id), nil
}

func (u *JobUsecases) Run(ctx context.Context, jobId string, arguments map[string]entities.Variable) (entities.Output, error) {
	job, err := u.repo.GetByID(jobId)
	if err != nil {
		return entities.Output{}, fmt.Errorf("failed to run job: %v", err)
	}

	jobSpace := make(map[string]interface{})

	for key, value := range arguments {
		jobSpace[key] = value.Value
	}

	for _, actionInfo := range job.Actions {
		action, err := u.aRepo.GetByID(ctx, string(actionInfo.Id))
		if err != nil {
			return job.Output.OnFailure, fmt.Errorf("failed to run job: %v", err)
		}

		for key, value := range action.Input {
			if value.Value == nil {
				action.Input[key] = entities.Variable{Value: jobSpace[key]}
			}
		}

		switch action.Type {
		case entities.Request:
			result, err := u.executeHTTPAction(*action)
			if err != nil {
				return job.Output.OnFailure, fmt.Errorf("failed to run job: %v", err)
			}

			for _, key := range actionInfo.Output {
				jobSpace[key] = result[key].Value
			}
		default:
			return job.Output.OnFailure, fmt.Errorf("failedt to run job: unknown action")
		}
	}
	return job.Output.OnSuccess, nil
}

func (u *JobUsecases) executeHTTPAction(action entities.Action) (map[string]entities.Variable, error) {
	client := &http.Client{}

	body := make(map[string]interface{})

	if len(action.Input) != 0 {
		for key, value := range action.Input {
			body[key] = value.Value
		}
	}
	b, err := json.Marshal(body)

	req, err := http.NewRequest(action.AdditionalParams["method"].(string), action.AdditionalParams["url"].(string), bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("failed to execute http action: %v", err)
	}

	req.Header.Set("Authorization", action.Input["token"].Value.(string))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute http action: %v", err)
	}
	defer resp.Body.Close()

	parsedResponse := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(parsedResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to execute http action: %v", err)
	}

	output := make(map[string]entities.Variable)
	for key, value := range parsedResponse {
		output[key] = entities.Variable{Type: "string", Value: value}
	}
	return output, nil
}
