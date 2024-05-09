package entities

type Pipeline struct {
	ID      PipelineID
	RootJob string
	jobs    map[string]Job
}

type PipelineID string
