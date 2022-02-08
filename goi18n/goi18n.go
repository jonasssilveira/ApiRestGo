package goi18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle
var messageError = &i18n.Message{ //1
	ID:  "error",
	One: "Ocorreu um erro, {{.Error}}",
}

func InitMessages() {
	bundle = i18n.NewBundle(language.BrazilianPortuguese)
	localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String())
}

func GetMessageError(erro string) string {
	translationOne, _ :=
		localizer.Localize(&i18n.LocalizeConfig{ //2
			DefaultMessage: messageError,
			TemplateData: map[string]interface{}{ //3
				"Error": erro,
			},
			PluralCount: 1, //4
		})
	return translationOne
}
