package errors

import (
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string `json:"field"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ValidationError struct {
	HttpError
	Errors []FieldError `json:"errors"`
}

func NewValidationError(path string, status int) *ValidationError {
	httpError := NewHttpError(
		"validation_error",
		"Ocorreu um erro na validação dos campos",
		"Um ou mais campos não foram preenchidos corretamente.",
		path,
		status,
	)
	return &ValidationError{
		HttpError: *httpError,
		Errors:    []FieldError{},
	}
}

func (e *ValidationError) Map(lng string, validationErrors validator.ValidationErrors, t *ut.Translator) *ValidationError {
	if t == nil {
		ptBR := pt_BR.New()
		uni := ut.New(ptBR, ptBR)
		utpt, _ := uni.GetTranslator("pt_BR")
		t = &utpt
	}
	for _, v := range validationErrors {
		field := v.Field()
		message := v.Translate(*t)
		code := v.Tag()

		e.Errors = append(e.Errors, FieldError{
			Field:   field,
			Code:    code,
			Message: message,
		})
	}

	return e
}
