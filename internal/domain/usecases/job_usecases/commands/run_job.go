package commands

import (
	"context"
	"fmt"

	"worker/internal/domain/entities"
)

type RunJobCommand struct {
}

func NewRunJobCommand() RunJobCommand {
	return RunJobCommand{}
}

func (c *RunJobCommand) Execute(ctx context.Context, job entities.Job) error {
	switch job.Type {
	case entities.HTTP:

	case entities.Condition:

	default:
		return fmt.Errorf("failed to execute command: unknown type")
	}
}

func (c *RunJobCommand) ExecuteConditionJob() (bool, error) {

}
