package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/shipt/shipt-tofu/http/server/response"
	"github.com/smwest87/dnd-golang/character"
	config "github.com/smwest87/dnd-golang/configuration"
)

var password = os.Getenv("DB_PASSWORD")

type endpointFunc func(r *http.Request) (int, []byte, error)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func GetCharacter(w http.ResponseWriter, r *http.Request) (int, []byte, error) {
	vars := mux.Vars(r)
	key := vars["id"]
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return 400, nil, err
	}

	queryCharacter := "SELECT * FROM dnd.dnd_characters WHERE id = $1"
	result := db.QueryRow(queryCharacter, key)
	var returnCharacter character.Character
	err = result.Scan(&returnCharacter.ID, &returnCharacter.Name, &returnCharacter.Class, &returnCharacter.Level, &returnCharacter.HitPointMaximum, &returnCharacter.Strength, &returnCharacter.Dexterity, &returnCharacter.Constitution, &returnCharacter.Wisdom, &returnCharacter.Intelligence, &returnCharacter.Charisma, &returnCharacter.Initiative, pq.Array(&returnCharacter.Modifiers))
	if err != nil {
		return 400, nil, err
	}

	json_returnCharacter, err := json.Marshal(returnCharacter)

	if err != nil {
		return 400, nil, err
	}

	return 200, json_returnCharacter, nil

}

func DeleteCharacter(w http.ResponseWriter, r *http.Request) (int, []byte, error) {
	vars := mux.Vars(r)
	key := vars["id"]
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return 400, nil, err
	}

	deleteCommand := "DELETE FROM dnd.dnd_characters WHERE id = $1"
	_, err = db.Exec(deleteCommand, key)
	if err != nil {
		return 400, nil, err
	}

	return 200, nil, nil
}

func CreateCharacter(w http.ResponseWriter, r *http.Request) (int, []byte, error) {
	returnCharacter := character.Character{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return 400, nil, err
	}
	err = json.Unmarshal(body, &returnCharacter)
	if err != nil {
		return 400, nil, err
	}

	hero, err := character.GenerateCharacter(returnCharacter.Name, returnCharacter.Class)

	if err != nil {
		return 400, nil, err
	}

	_, err = character.InsertCharacter(*hero)

	if err != nil {
		return 400, nil, err
	}

	json_hero, err := json.Marshal(hero)
	if err != nil {
		return 400, nil, err
	}

	return 200, json_hero, nil

}

func UpdateCharacter(w http.ResponseWriter, r *http.Request) (int, []byte, error) {
	returnCharacter := character.Character{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return 400, nil, err
	}
	err = json.Unmarshal(body, &returnCharacter)
	if err != nil {
		return 400, nil, err
	}

	fields := reflect.TypeOf(returnCharacter)
	values := reflect.ValueOf(returnCharacter)

	totalFields := fields.NumField()

	for i := 0; i < totalFields; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		fmt.Print("Type", field.Type, ",", field.Name, "=", value, "\n")
	}

	return 200, nil, nil

}

func ResponseWrapper(f endpointFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode, payload, err := f(r)
		switch {
		case err != nil:
			// After launch, consider warnings for non 5xx errors
			response.Error(
				w,
				statusCode,
				err.Error(),
				err,
				nil,
			)
		default:
			response.WithPayload(
				w,
				statusCode,
				payload,
				nil,
			)
		}
	}
}
