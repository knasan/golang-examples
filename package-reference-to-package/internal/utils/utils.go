package utils

import (
	"fmt"

	"github.com/knasan/golang-examples/package-reference-to-package/internal/config"
)

var app *config.AppConfig

func NewAppConfig(a *config.AppConfig) {
	app = a
}

func ShowDebug() {
	fmt.Println(app.Debug)
}
