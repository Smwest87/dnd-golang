package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/smwest87/dnd-golang/character"
	config "github.com/smwest87/dnd-golang/configuration"
)

var password = os.Getenv("DB_PASSWORD")

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func GetCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	queryCharacter := "SELECT * FROM dnd.dnd_characters WHERE id = $1"
	result := db.QueryRow(queryCharacter, key)
	var returnCharacter character.Character
	err = result.Scan(&returnCharacter.ID, &returnCharacter.Name, &returnCharacter.Class, &returnCharacter.Level, &returnCharacter.HitPointMaximum, &returnCharacter.Strength, &returnCharacter.Dexterity, &returnCharacter.Constitution, &returnCharacter.Wisdom, &returnCharacter.Intelligence, &returnCharacter.Charisma, &returnCharacter.Initiative, &returnCharacter.Modifiers)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(returnCharacter)

}

func DeleteCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal()
	}

	deleteCommand := "DELETE FROM dnd.dnd_characters WHERE id = $1"
	_, err = db.Exec(deleteCommand, key)
	if err != nil {
		log.Fatal()
	}
}

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	returnCharacter := character.Character{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	err = json.Unmarshal(body, &returnCharacter)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(returnCharacter)

	hero, err := character.GenerateCharacter(returnCharacter.Name, returnCharacter.Class)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(hero)

}

func UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	returnCharacter := character.Character{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	err = json.Unmarshal(body, &returnCharacter)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

}
