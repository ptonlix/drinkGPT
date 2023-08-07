package prompt

import (
	"embed"

	"github.com/ptonlix/drinkGPT/common/util"
)

//go:embed templates/*
var templatesFS embed.FS

const (
	DrinkChatTemplate = "drinkchat.tmpl"
)

// Initializes the prompt package by loading the templates from the embedded file system.
func init() {
	if err := util.LoadTemplates(templatesFS); err != nil {
		panic(err)
	}
}
