package entities

type Job struct {
	Id      JobID
	Name    string
	Actions []ActionInformation
	Output  JobOutput
}

type JobID string
type JobType int

type JobOutput struct {
	Ask       Output
	OnSuccess Output
	OnFailure Output
}

type Output struct {
	Message  string
	Type     OutputType
	Variable map[string]Variable
}

type OutputType int

const (
	Ask OutputType = iota + 1
	Success
	Failure
)

type ActionInformation struct {
	Id     ActionID
	Output []string
}

type JobRepository interface {
	Create(job Job) error
	GetByID(id string) (Job, error)
	NextID() JobID
}

func NewJob(id JobID, actions []ActionInformation, output JobOutput) (*Job, error) {
	return &Job{
		Id:      id,
		Actions: actions,
		Output:  output,
	}, nil
}
