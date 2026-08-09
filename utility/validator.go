package utility

import (
	"errors"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func ValidatorG(s interface{}) (bool, string) {

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(s)
	if err == nil {
		return true, ""
	}

	english := en.New()

	uni := ut.New(english, english)

	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, trans)

	var validateErrors validator.ValidationErrors

	if !errors.As(err, &validateErrors) {
		return false, err.Error()
	} else {
		var errorList []string
		for _, fieldError := range validateErrors {

			if fieldError.Field() == "CountryId" {
				errorList = append(errorList, "Country not found!")
				continue
			}

			errorList = append(errorList, fieldError.Translate(trans))
		}
		return false, strings.Join(errorList, " , ")
	}
}
