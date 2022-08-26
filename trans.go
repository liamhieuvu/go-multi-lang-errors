package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var translations = map[string]map[string]string{
	"en": {
		"required": "{0} is a required field.",
		"email":    "{0} is invalid.",
		"gte":      "{0} must be {1} or greater.",
		"lte":      "{0} must be {1} or smaller.",
	},
	"vi": {
		"required": "{0} là trường bắt buộc.",
		"email":    "{0} không hợp lệ.",
		"gte":      "{0} phải bằng hoặc lớn hơn {1}.",
		"lte":      "{0} phải bằng hoặc nhỏ hơn {1}.",
	},
}

var dicts = map[string]map[string]string{
	"vi": {
		"Name": "Tên",
		"Age":  "Tuổi",
	},
}

func translationFunc(t ut.Translator, fe validator.FieldError) string {
	field, err := t.T(fe.Field())
	if err != nil {
		field = fe.Field()
	}
	msg, err := t.T(fe.Tag(), field, fe.Param())
	if err != nil {
		return fe.Error()
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

func getErrMsg(errs validator.ValidationErrorsTranslations) string {
	messages := make([]string, 0, len(errs))
	for _, v := range errs {
		messages = append(messages, v)
	}
	return strings.Join(messages, " ")
}
