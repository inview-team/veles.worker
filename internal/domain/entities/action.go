package entities

type Action struct {
	Id               ActionID
	Type             ActionType
	Arguments        map[string]Variable
	AdditionalParams map[string]interface{}
}

type ActionID string

type ActionType int

const (
	HTTP ActionType = iota + 1
	Compare
)

type Variable struct {
	Value interface{}
	Type  string
}

type Parameters struct {
	HTTPRequest ``
}

type HTTPRequest struct {
	Url     string
	Method  string
	Headers map[string]string
	Body    map[string]interface{}
}

type CompareOperations int

const (
	Equal CompareOperations = iota + 1
	More
)

type ActionRepository interface {
	Create(action Action) error
	GetByID(id string) (Action, error)
	NextID() ActionID
}

func NewAction(id ActionID, actionType ActionType, arguments map[string]Variable) (*Action, error) {
	return &Action{
		Id:        id,
		Type:      actionType,
		Arguments: arguments,
	}, nil
}
