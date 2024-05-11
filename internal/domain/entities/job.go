package entities

type Job struct {
	Id       JobID
	ActionID ActionID
	Output   []string
}

type JobID string

type JobRepository interface {
	Create(job Job) error
	GetByID(id string) (Job, error)
	NextID() JobID
}

func NewJob(id JobID, actionID ActionID, output []string) (*Job, error) {
	return &Job{
		Id:       id,
		ActionID: actionID,
		Output:   output,
	}, nil
}
