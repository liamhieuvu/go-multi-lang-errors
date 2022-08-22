package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/vi"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type user struct {
	Name  string `validate:"required"`
	Age   uint8  `validate:"gte=10,lte=90"`
	Email string `validate:"required,email"`
}

var (
	validate *validator.Validate
	utrans   *ut.UniversalTranslator
)

func main() {
	setupTrans()
	r := gin.Default()
	r.POST("/default/users", handleDefault)
	r.POST("/trans/users", handleTranslation)
	_ = r.Run()
}

func handleDefault(c *gin.Context) {
	var u user
	if err := c.ShouldBindJSON(&u); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := validate.Struct(u); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}

func handleTranslation(c *gin.Context) {
	var u user
	if err := c.ShouldBindJSON(&u); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := validate.Struct(u); err != nil {
		transErrs := err.(validator.ValidationErrors).Translate(getTransFromParam(c))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": getErrMsg(transErrs)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}

func setupTrans() {
	validate = validator.New()
	enLocale := en.New()
	utrans = ut.New(enLocale, enLocale, vi.New())

	for locale, trans := range translations {
		t, _ := utrans.FindTranslator(locale)
		for tag, translation := range trans {
			_ = validate.RegisterTranslation(tag, t, getRegisterFunc(tag, translation), translationFunc)
		}
	}
}
