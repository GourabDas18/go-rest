package utility

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func ValidatorG() (*validator.Validate, *ut.Translator) {

	validator := validator.New(validator.WithRequiredStructEnabled())

	enLocal := en.New()
	ut, _ := ut.New(enLocal).GetTranslator("en")
	return validator, &ut
}
