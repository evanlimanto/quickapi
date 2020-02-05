package main

import (
	"github.com/evanlimanto/quickapi/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.initializeRoutes()
	app.Run()
}
