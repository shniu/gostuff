package errors

import "errors"

var (
	ErrNotFound        = New("kvs: not found")
	EngineTypeNotFound = New("Engine type: not found")
)

func New(text string) error {
	return errors.New(text)
}
