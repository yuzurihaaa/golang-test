package routes

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/yuzuriha/restapi/service"
)

func HandleRoute() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user", service.RegisterUser).Methods("POST")

	return router
}
