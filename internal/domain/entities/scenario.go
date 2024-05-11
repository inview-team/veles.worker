package entities

import "fmt"

type Scenario struct {
	ID          ScenarioID
	RootJobID   JobID
	jobs        []JobID
	JobSequence map[JobID]JobID
}

type ScenarioID string

type ScenarioRepository interface {
	Create(scenario Scenario) error
	GetById(Id string) (Scenario, error)
	Delete(Id string) (Scenario, error)
	NextID() ScenarioID
}

func NewScenario(id ScenarioID, rootJobID JobID, jobs []JobID) (*Scenario, error) {
	return &Scenario{
		ID:        id,
		RootJobID: rootJobID,
		jobs:      jobs,
	}, nil
}

func (s *Scenario) GetNextJob(id JobID) (JobID, error) {
	value, ok := s.JobSequence[id]
	if !ok {
		return "", fmt.Errorf("failed to get next job: job not found")
	}

	if string(value) == "-1" {
		return "", fmt.Errorf("failed to get next job: job don't have next job")
	}
	return value, nil
}
