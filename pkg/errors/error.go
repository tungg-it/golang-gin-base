package errorException

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type BadRequestResponse struct {
	Message    string `json:"message" example:"token_invalid"`
	StatusCode int    `json:"status_code" example:"400"`
}

type InternalServerResponse struct {
	Message    string `json:"message" example:"internal_server_error"`
	StatusCode int    `json:"status_code" example:"500"`
}

type NotFoundResponse struct {
	Message    string `json:"message" example:"user_not_found"`
	StatusCode int    `json:"status_code" example:"404"`
}

type Http struct {
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"status_code"`
}

func (e Http) Error() string {
	return fmt.Sprintf("message: %s", e.Message)
}

func NewHttpError(statusCode int, message string) Http {
	return Http{
		Message:    message,
		StatusCode: statusCode,
	}
}

func ValidateStruct(input interface{}) *string {
	var errors []string

	validation := validator.New()

	err := validation.Struct(input)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			customErr := validateErrorMessage(e)
			errors = append(errors, customErr)
		}
	}

	if len(errors) > 0 {
		message := strings.ReplaceAll(strings.ToLower(errors[0]), " ", "_")
		return &message
	}

	return nil

}

func validateErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("Field %s is required", err.Field())
	case "email":
		return fmt.Sprintf("Field %s wrong email format", err.Field())
	case "max":
		return fmt.Sprintf("Field %s max %s charset", err.Field(), err.Param())
	default:
		return fmt.Sprintf("Validation error on field %s with tag %s", err.Field(), err.Tag())
	}
}
