package main

import (
	"github.com/knasan/golang-examples/package-reference-to-package/internal/config"
	"github.com/knasan/golang-examples/package-reference-to-package/internal/utils"
)

var app config.AppConfig

func main() {
	app.Debug = true
	utils.NewAppConfig(&app)
	utils.ShowDebug()
	app.Debug = false
	utils.ShowDebug()
}
