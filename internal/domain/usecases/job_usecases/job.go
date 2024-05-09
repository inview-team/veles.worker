package job_usecases

import "worker/internal/domain/usecases/job_usecases/commands"

type JobUsecases struct {
	Commands
}

type Commands struct {
	RunJob commands.RunJobCommand
}

func NewJobUsecases() JobUsecases {
	return JobUsecases{
		Commands: Commands{
			RunJob: commands.NewRunJobCommand(),
		},
	}
}
