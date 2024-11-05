package routes

import (
	"database/sql"
	"timesheet-app/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, db *sql.DB) {

	r.HandleFunc("/login", controllers.LoginHandler(db)).Methods("POST")
}
