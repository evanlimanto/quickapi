package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	DB     *sql.DB
	Router *mux.Router
}

func (app *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/v0/balance/get", app.GetBalance).Methods("GET")
}
