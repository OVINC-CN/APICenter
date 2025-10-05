package i18n

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle
var DefaultLang = language.SimplifiedChinese
var localeFilePath = "locales"

func init() {
	// init
	bundle = i18n.NewBundle(DefaultLang)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	if err := filepath.WalkDir(
		localeFilePath,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				if _, err := bundle.LoadMessageFile(path); err != nil {
					log.Printf("[I18n] load locales %s fail: %s\n", path, err.Error())
				}
				log.Printf("[I18n] load locales %s success\n", path)
			}
			return nil
		},
	); err != nil {
		log.Printf("[I18n] read locales dir %s fail: %s\n", localeFilePath, err.Error())
	}
	log.Printf("[I18n] init success\n")
}

func getLocalizer(lang string) *i18n.Localizer {
	if lang == "" {
		return getLocalizer(DefaultLang.String())
	}
	return i18n.NewLocalizer(bundle, lang)
}

func Translate(lang string, messageID string, data map[string]string) string {
	msg, err := getLocalizer(lang).Localize(
		&i18n.LocalizeConfig{
			MessageID:      messageID,
			TemplateData:   data,
			DefaultMessage: &i18n.Message{ID: messageID, Other: messageID},
		},
	)
	if err != nil {
		tmpl := template.New("msg").Option("missingkey=zero")
		t, err := tmpl.Parse(messageID)
		if err != nil {
			return messageID
		}
		var buf strings.Builder
		if err := t.Execute(&buf, data); err != nil {
			return messageID
		}
		return buf.String()
	}
	return msg
}
