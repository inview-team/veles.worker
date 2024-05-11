package sessions

import (
	"context"

	"worker/internal/domain/entities"
	"worker/internal/domain/usecases/dialog_usecases"
)

type App struct {
	Session dialog_usecases.DialogUsecases
}

func NewApp(ctx context.Context, sessionRepo entities.DialogRepository) (*App, error) {
	session := dialog_usecases.New(sessionRepo)
	app := App{
		Session: *session,
	}
	return app, nil
}
