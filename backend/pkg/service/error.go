package service

import (
	"fmt"
)

type DatabaseError struct {
	operation, model string
	content error
}
func (err DatabaseError) Error() string { return fmt.Sprintf("%v %v: %v", err.operation, err.model, err.content) }

type ServiceNotFoundError string
func (err ServiceNotFoundError) Error() string { return fmt.Sprintf("service not found: %v", string(err)) }

type ServiceAlreadyExistError string
func (err ServiceAlreadyExistError) Error() string { return fmt.Sprintf("service already exist: %v", string(err)) }

type InvalidUrlError string
func (err InvalidUrlError) Error() string { return fmt.Sprintf("invalid url: %v", string(err)) }