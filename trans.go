package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// locale: tag: translation
var translations = map[string]map[string]string{
	"en": {
		"required": "{0} is a required field",
		"email":    "{0} is invalid",
	},
	"vi": {
		"required": "{0} là trường bắt buộc",
		"email":    "{0} không hợp lệ",
	},
}

func getRegisterFunc(tag, translation string) validator.RegisterTranslationsFunc {
	return func(t ut.Translator) error {
		return t.Add(tag, translation, false)
	}
}

func translationFunc(t ut.Translator, fe validator.FieldError) string {
	msg, err := t.T(fe.Tag(), fe.Field())
	if err != nil {
		fmt.Printf("warning: error translating FieldError: %#v\n", fe)
		return fe.(error).Error()
	}
	return msg
}

func getTransFromParam(c *gin.Context) ut.Translator {
	t, found := utrans.GetTranslator(c.Query("locale"))
	if !found {
		t, _ = utrans.GetTranslator("en")
	}
	return t
}
