package goi18n

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func InitMessages() *i18n.Localizer {
	bundle = i18n.NewBundle(language.BrazilianPortuguese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("ressources/ptBr.json")
	return i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String())
}

//func GetMessage(idMessage string) string {
//	localizeConfig := i18n.LocalizeConfig{
//		MessageID: idMessage,
//	}
//	localizationUsingJson, _ := localizer.Localize(&localizeConfig)
//	return localizationUsingJson
//}
