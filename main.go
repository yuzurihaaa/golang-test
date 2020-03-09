package main

//go:generate sqlboiler --wipe psql

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/yuzuriha/restapi/routes"
	"github.com/yuzuriha/restapi/util"
	"net/http"
)

func main() {

	fmt.Println("connected")

	router := routes.HandleRoute()
	err := http.ListenAndServe(":8000", router)

	util.DieIf(err)
}
