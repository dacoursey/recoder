package main

import (
	"database/sql"
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// App for stuff.
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize for stuff.
func (a *App) Initialize(user, password, dbname string) {
	//connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	// var err error
	// a.DB, err = sql.Open("postgres", connectionString)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	store.Options = &sessions.Options{
		Domain: "localhost",
		Path:   "/",
		MaxAge: 3600 * 8, // 8 hours
		// HttpOnly: true, // disable for this demo
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

// Build up our routing table.
func (a *App) initializeRoutes() {
	// AuthZ and AuthN
	a.Router.HandleFunc("/login", a.processLogin).Methods("POST")
	a.Router.HandleFunc("/logout", a.processLogout).Methods("GET")
	// Static Files
	a.Router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(rice.MustFindBox("public").HTTPBox())))
	// Coders
	a.Router.HandleFunc("/recode", a.recode).Methods("POST")
}

// Run for stuff.
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
