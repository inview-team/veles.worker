package job_usecases

import (
	"fmt"

	"worker/internal/domain/entities"
)

type JobUsecases struct {
	repo entities.JobRepository
}

func New(repo entities.JobRepository) (*JobUsecases, error) {
	return &JobUsecases{
		repo: repo,
	}, nil
}

func (u *JobUsecases) Create(actionID string, output []string) (string, error) {
	job, err := entities.NewJob(u.repo.NextID(), entities.ActionID(actionID), output)
	if err != nil {
		return "", fmt.Errorf("failed to create job: %v", err)
	}

	err = u.repo.Create(*job)
	if err != nil {
		return "", fmt.Errorf("failed to create job: %v", err)
	}
	return string(job.Id), nil
}
