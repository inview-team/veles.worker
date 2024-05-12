package executor

import (
	"context"

	"github.com/inview-team/veles.worker/pkg/domain/usecases/action_usecases"
	"github.com/inview-team/veles.worker/pkg/domain/usecases/job_usecases"
	"github.com/inview-team/veles.worker/pkg/infrastructure/mongodb"
	"github.com/inview-team/veles.worker/pkg/infrastructure/mongodb/action_repository"
	"github.com/inview-team/veles.worker/pkg/infrastructure/mongodb/job_repository"
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
