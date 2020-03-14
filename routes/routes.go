package routes

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/yuzuriha/restapi/service"
)

func HandleRoute() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user", service.RegisterUser).Methods("POST")
	router.HandleFunc("/api/user", service.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/", service.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/user/", service.FindUser).Methods("GET")

	return router
}
