package main

import (
	"github.com/cherryReptile/WS-APP/internal/app"
)

func main() {
	app := app.NewApp()
	app.Run("80")
}
