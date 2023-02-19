package utils

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func GetString(val string, lang language.Tag) string {
	bundle := i18n.NewBundle(lang)
	bundle.MustLoadMessageFile("./languages/active." + lang.String() + ".json")
	localizer := i18n.NewLocalizer(bundle, lang.String())
	msg := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: val,
		},
	})
	return msg
}
