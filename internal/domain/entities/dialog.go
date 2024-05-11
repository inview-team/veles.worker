package entities

type Dialog struct {
	ID           DialogID
	ClientID     ClientID
	ScenarioID   ScenarioID
	CurrentJobID JobID
	Storage      map[string]interface{}
}

type DialogID string
type ClientID string

type DialogRepository interface {
	Create(dialog Dialog) error
	GetById(Id string) (Dialog, error)
	Delete(Id string) (Dialog, error)
	Update(dialog Dialog) error
	NextID() DialogID
}

type DialogResult struct {
	Argument        map[string]Variable
	Output          map[string]Variable
	IsNeedArguments bool
}

func NewDialog(id DialogID, clientID ClientID) (*Dialog, error) {
	return &Dialog{
		ID:           id,
		ClientID:     clientID,
		ScenarioID:   "",
		CurrentJobID: "",
		Storage:      make(map[string]interface{}),
	}, nil
}

func (s *Dialog) AddVariable(name string, value interface{}) {
	s.Storage[name] = value
}

func (s *Dialog) GetVariable(name string) (interface{}, bool) {
	val, ok := s.Storage[name]
	return val, ok
}
