package controller

import (
	"reflect"
	"sync"

	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/pt"
	"github.com/gofiber/fiber/v2"
)

var (
	v      *validator.Validate
	once   sync.Once
	status = fiber.StatusBadRequest
	trans  *ut.Translator
)

func getValidator() (*validator.Validate, *ut.Translator) {
	once.Do(func() {
		v = validator.New()
		// Configurar tradução para pt-BR
		ptBR := pt_BR.New()
		uni := ut.New(ptBR, ptBR)
		t, _ := uni.GetTranslator("pt_BR")
		trans = &t
		// Registrar tradução de mensagens
		_ = pt.RegisterDefaultTranslations(v, t)

		// Registra JSON como Tag
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("json")
		})
	})

	return v, trans
}
