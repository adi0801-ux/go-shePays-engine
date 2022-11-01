package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"shepays/constants"
)

type IError struct {
	Type    string `json:"type"`
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

func ValidateRequest[C any](s *HTTPServer, c *fiber.Ctx, requestBody *C) ([]IError, error) {

	if err := c.BodyParser(&requestBody); err != nil {
		if err.Error() == constants.UnprocessableEntity {
			customError := IError{
				Type:    "Invalid json",
				Message: "Json cannot be parsed",
			}
			return []IError{customError}, fmt.Errorf(customError.Message)
		}
		te, _ := err.(*json.UnmarshalTypeError)

		customError := IError{
			Type:    "DataType Error",
			Field:   te.Field,
			Tag:     te.Type.String(),
			Value:   te.Value,
			Message: "Change the data type of the field to " + te.Type.String(),
		}
		return []IError{customError}, fmt.Errorf(customError.Message)
	}

	if err := s.validator.Struct(requestBody); err != nil {
		var errors []IError
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			el.Type = "Validation Error"
			el.Message = "Please refer docs for providing valid field value"

			errors = append(errors, el)
		}

		return errors, fmt.Errorf(constants.ValidationError)
	}

	return []IError{}, nil
}
