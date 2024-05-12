package sessions

import (
	"context"

	"worker/internal/domain/usecases/action_usecases"
	"worker/internal/domain/usecases/dialog_usecases"
	"worker/internal/domain/usecases/job_usecases"
)

type App struct {
	Job    job_usecases.JobUsecases
	Action action_usecases.ActionUsecases
}

func NewApp(ctx context.Context) (*App, error) {
	session := dialog_usecases.New(sessionRepo)
	app := App{
		Session: *session,
	}
	return app, nil
}
