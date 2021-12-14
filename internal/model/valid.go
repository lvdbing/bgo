package model

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"

	validator "github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func ValidAndBind(c *gin.Context, value interface{}) (bool, ValidErrors) {
	var validErrors ValidErrors
	err := c.ShouldBind(value)
	if err != nil {
		cv := c.Value("trans")
		trans, _ := cv.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, validErrors
		}
		for k, v := range verrs.Translate(trans) {
			validErrors = append(validErrors, &ValidError{
				Key:     k,
				Message: v,
			})
		}
		return false, validErrors
	}

	return true, nil
}
