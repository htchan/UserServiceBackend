package user

import (
	"fmt"
)

type InvalidParamsError string
func (err InvalidParamsError) Error() string { return fmt.Sprintf("invalid %v", string(err)) }

type DuplicatedUserError struct{}
func (err DuplicatedUserError) Error() string { return "user already exist" }

type IncorrectParamsError string
func (err IncorrectParamsError) Error() string { return fmt.Sprintf("incorrect %v", string(err)) }

type DatabaseError struct {
	operation, model string
	content error
}
func (err DatabaseError) Error() string { return fmt.Sprintf("%v %v: %v", err.operation, err.model, err.content) }