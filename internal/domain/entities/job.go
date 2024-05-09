package entities

type Job struct {
	Id             JobID
	Type           JobType
	Next           NextJob
	InputTemplate  string
	OutputTemplate string
}

type JobType int
type JobID string

type NextJob struct {
	OnSuccess JobID
	OnFailure JobID
}

const (
	HTTP JobType = iota + 1
	Condition
)
