package action_usecases

import (
	"fmt"

	"worker/internal/domain/entities"
)

type ActionUsecases struct {
	repo entities.ActionRepository
}

func New(repo entities.ActionRepository) (*ActionUsecases, error) {
	return &ActionUsecases{repo: repo}, nil
}

func (u *ActionUsecases) Register(actionType entities.ActionType, arguments map[string]entities.Variable, params map[string]interface{}) (string, error) {
	action, err := entities.NewAction(u.repo.NextID(), actionType, arguments, params)
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
