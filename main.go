package main

import (
	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// import "godotenv" here ...
	"github.com/joho/godotenv"
)

func main() {

	// Init godotenv here ...
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	// biar folder uploads bisa diakses
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	// Initialization "uploads" folder to public here ...

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
