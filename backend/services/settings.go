package services

import (
	"my-app/backend/app"
	"my-app/backend/model"
)

func SaveColorThemeOption(theme string) error {
	option := model.MyOption{
		Name: app.CfgColorTheme,
	}
	result := option.Update(theme)
	return result.Error
}

func SaveDisplayLanguageOption(lang string) error {
	option := model.MyOption{
		Name: app.CfgDisplayLanguage,
	}
	result := option.Update(lang)
	return result.Error
}
