package util

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/go-playground/validator.v10"
	"net/http"
)

type databaseConfig struct {
	Dbname  string
	User    string
	Pass    string
	Host    string
	Sslmode string
}

type Config struct {
	Psql databaseConfig
}

func GetDatabase() *sql.DB {

	var config Config

	if _, err := toml.DecodeFile("./sqlboiler.toml", &config); err != nil {
		fmt.Println(err)
		return nil
	}

	connString := fmt.Sprintf(`dbname=%s host=%s user=%s password=%s sslmode=%s`, config.Psql.Dbname, config.Psql.Host, config.Psql.User, config.Psql.Pass, config.Psql.Sslmode)

	db, err := sql.Open("postgres", connString)
	DieIf(err)

	err = db.Ping()
	DieIf(err)

	return db
}

func DieIf(err error) {
	if err != nil {
		panic(err)
	}
}

func Validate(w http.ResponseWriter, v interface{}) error {
	validate := validator.New()
	if err := validate.Struct(v); err != nil {
		validationError := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			validationError[e.Field()] = e.ActualTag()
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(validationError)
		return err
	}

	return nil
}

func VerifyAndDecode(w http.ResponseWriter, request *http.Request, v interface{}) error {
	if err := json.NewDecoder(request.Body).Decode(v); err != nil {
		return errors.New("ERROR Decode")
	}

	err := Validate(w, v)

	return err
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}
