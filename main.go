package main

import (
	"github.com/evanlimanto/quickapi/src/api"
)

func main() {
	app := &api.App{}
	app.Initialize("username", "password", "dbname")
	app.Run()
}
