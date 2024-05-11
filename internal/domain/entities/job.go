package entities

type Job struct {
	Id        JobID
	Arguments []Variable
	Actions   []ActionID
	Output    JobOutput
}

type JobID string

type JobType int

type JobOutput struct {
	Message  string
	Type     OutputType
	Variable Variable
}

type Variable struct {
	Value interface{}
	Type  string
}

type OutputType int

const (
	Ask OutputType = iota + 1
	Success
	Failure
)

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
