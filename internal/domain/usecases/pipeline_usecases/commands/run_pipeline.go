package commands

import "context"

type RunPipelineCommand struct {
}

func NewRunPipelineCommand() RunPipelineCommand {
	return RunPipelineCommand{}
}

func (c *RunPipelineCommand) Execute(ctx context.Context) (string, error) {

}
