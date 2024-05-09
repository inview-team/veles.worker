package pipeline_usecases

import "worker/internal/domain/usecases/pipeline_usecases/commands"

type PipelineUsecases struct {
	Commands
}

type Commands struct {
	RunPipeline commands.RunPipelineCommand
}

func NewPipelineUsecases() PipelineUsecases {
	return PipelineUsecases{
		Commands{
			RunPipeline: commands.NewRunPipelineCommand(),
		},
	}
}
