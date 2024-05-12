package executor

import (
	"context"

	"worker/internal/domain/usecases/action_usecases"
	"worker/internal/domain/usecases/job_usecases"
	"worker/internal/infrastructure/mongodb"
	"worker/internal/infrastructure/mongodb/action_repository"
	"worker/internal/infrastructure/mongodb/job_repository"
)

type App struct {
	JobUseCases    job_usecases.JobUsecases
	ActionUseCases action_usecases.ActionUsecases
}

func NewApp(ctx context.Context, cfg mongodb.Config) (*App, error) {
	c, err := mongodb.NewClient(ctx, cfg)
	if err != nil {
		return nil, err
	}

	actRepo := action_repository.NewActionRepository(c)
	jobRepo := job_repository.NewJobRepository(c)

	app := App{
		JobUseCases:    job_usecases.New(jobRepo, actRepo),
		ActionUseCases: action_usecases.New(actRepo),
	}
	return &app, nil
}
