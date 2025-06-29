package request

import (
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	// contextWrapperService defines the methods for a context wrapper.
	contextWrapperService interface {
		Bind(data any) error
	}
	// contextWrapper is a wrapper around echo.Context that provides additional functionality.
	contextWrapper struct {
		Context   echo.Context
		Validator *validator.Validate
	}
)


// NewContextWrapper creates a new context wrapper for the provided echo context.
func NewContextWrapper(ctx echo.Context) contextWrapperService {
	return &contextWrapper{
		Context:   ctx,
		Validator: validator.New(),
	}
}
// Bind binds the request data to the provided struct and validates it.
func (c *contextWrapper) Bind(data any) error {
	if err := c.Context.Bind(data); err != nil {
		log.Printf("error binding data: %v", err)
		return err
	}

	if err := c.Validator.Struct(data); err != nil {
		log.Printf("validation error: %v", err)
		return err
	}

	return nil
}