package todos

import "errors"

type TodoIdNotExistsError struct {
}

func (m *TodoIdNotExistsError) Error() string {
	return "todo ID not exists"
}

var (
	LinkToInvalidTagIDError = errors.New("link to tagID that is invalid")
)
