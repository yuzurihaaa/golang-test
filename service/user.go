package service

import (
	"context"
	"encoding/json"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/yuzuriha/restapi/models"
	"github.com/yuzuriha/restapi/util"
	"net/http"
	"strings"
)

type Register struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

func RegisterUser(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body Register

	if err := util.VerifyAndDecode(w, request, &body); err != nil {
		return
	}

	user := &models.User{Name: body.FirstName + " " + body.LastName}

	if err := user.Insert(context.Background(), util.GetDatabase(), boil.Infer()); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(util.Response{Message: "Fail Insert user to database"})
		return
	}

	_ = json.NewEncoder(w).Encode(util.Response{Message: "Create Success", Data: user})
}

type Update struct {
	Id        int    `json:"id" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

func UpdateUser(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body Update

	if err := util.VerifyAndDecode(w, request, &body); err != nil {
		return
	}

	user, userError := models.FindUser(context.Background(), util.GetDatabase(), body.Id)

	if userError != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(util.Response{Message: "Fail to get this user"})
		return
	}

	if body.FirstName != "" {
		user.Name = body.FirstName + " " + strings.Fields(user.Name)[1]
	}

	if body.LastName != "" {
		user.Name = strings.Fields(user.Name)[0] + " " + body.LastName
	}

	_, err := user.Update(context.Background(), util.GetDatabase(), boil.Infer())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(util.Response{Message: "Fail to update user"})
		return
	}

	_ = json.NewEncoder(w).Encode(util.Response{Message: "Update Success", Data: user})
}
