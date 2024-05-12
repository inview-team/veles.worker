package action_usecases

import (
	"context"
	"fmt"

	"github.com/inview-team/veles.worker/pkg/domain/entities"
)

type ActionUsecases struct {
	repo entities.ActionRepository
}

func New(repo entities.ActionRepository) ActionUsecases {
	return ActionUsecases{repo: repo}
}

func (u *ActionUsecases) Register(ctx context.Context, actionName string, actionType entities.ActionType, arguments map[string]entities.Variable, output []string, params map[string]interface{}) (string, error) {
	action, err := entities.NewAction(u.repo.NextID(ctx), actionName, actionType, arguments, output, params)
	if err != nil {
		return "", fmt.Errorf("failed to register action: %v", err)
	}

	err = u.repo.Create(ctx, action)
	if err != nil {
		return "", fmt.Errorf("failed to register action: %v", err)
	}
	return string(action.Id), err
}

func (u *ActionUsecases) GetByID(ctx context.Context, actionId string) (*entities.Action, error) {
	return u.repo.GetByID(ctx, actionId)
}
