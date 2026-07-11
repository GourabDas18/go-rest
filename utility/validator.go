package utility

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
)

func ValidatorG() (*validator.Validate, ut.Translator) {

	validator := validator.New(validator.WithRequiredStructEnabled())

	ut := ut.New().GetTranslator("en")
	enG := en.RegisterDefaultTranslations(validator, ut)

}
