package dialog_usecases

import (
	"fmt"

	"worker/internal/domain/entities"
)

type DialogUsecases struct {
	repo   entities.DialogRepository
	scRepo entities.ScenarioRepository
	jRepo  entities.JobRepository
	aRepo  entities.ActionRepository
}

func New(repo entities.DialogRepository, aRepo entities.ActionRepository, jRepo entities.JobRepository, scRepo entities.ScenarioRepository) *DialogUsecases {
	return &DialogUsecases{
		repo:   repo,
		aRepo:  aRepo,
		jRepo:  jRepo,
		scRepo: scRepo,
	}
}

func (u *DialogUsecases) Create(clientID string) (string, error) {
	session, err := entities.NewDialog(u.repo.NextID(), entities.ClientID(clientID))
	if err != nil {
		return "", fmt.Errorf("failed to create sessions: %v", err)
	}

	err = u.repo.Create(*session)
	if err != nil {
		return "", fmt.Errorf("failed to create sessions: %v", err)
	}
	return string(session.ID), nil
}

func (u *DialogUsecases) StartScenario(dialogID string, scenarioID string) error {
	dialog, err := u.repo.GetById(dialogID)
	if err != nil {
		return fmt.Errorf("failed start scenario: %v", err)
	}

	scenario, err := u.scRepo.GetById(scenarioID)
	if err != nil {
		return fmt.Errorf("failed start scenario: %v", err)
	}

	dialog.ScenarioID = entities.ScenarioID(scenarioID)
	dialog.CurrentJobID = scenario.RootJobID
	err = u.repo.Update(dialog)
	if err != nil {
		return fmt.Errorf("failed start scenario: %v", err)
	}
	return nil
}

func (u *DialogUsecases) Move(dialogId string) (string, error) {
	dialog, err := u.repo.GetById(dialogId)
	if err != nil {
		return "", fmt.Errorf("failed move throw scenario: %v", err)
	}

	currentScenario, err := u.scRepo.GetById(string(dialog.ScenarioID))
	if err != nil {
		return "", fmt.Errorf("failed move throw scenario: %v", err)
	}

	currentJob, err := u.jRepo.GetByID(string(dialog.CurrentJobID))

	nextJobId, err := currentScenario.GetNextJob(currentJob.Id)
	if err != nil {
		return "", fmt.Errorf("failed move throw scenario: %v", err)
	}
	dialog.CurrentJobID = nextJobId
	err = u.repo.Update(dialog)
	if err != nil {
		return fmt.Errorf("failed move throw scenario: %v", err)
	}
}
