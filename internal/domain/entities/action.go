package entities

type Action struct {
	Id               ActionID
	Name             string
	Type             ActionType
	Input            map[string]Variable
	Output           map[string]Variable
	AdditionalParams map[string]interface{}
}

type ActionID string

type ActionType int

const (
	Request ActionType = iota + 1
	Validate
)

type Variable struct {
	Value interface{}
	Type  string
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

func NewAction(id ActionID, actionType ActionType, arguments map[string]Variable, params map[string]interface{}) (*Action, error) {
	return &Action{
		Id:               id,
		Type:             actionType,
		AdditionalParams: params,
	}, nil
}
