package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/smwest87/dnd-golang/character"
)

var password = os.Getenv("DB_PASSWORD")

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func GetCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", character.Host, character.Port, character.User, password, character.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal()
	}

	queryCharacter := "SELECT * FROM dnd.dnd_characters WHERE id = $1"

	result := db.QueryRow(queryCharacter, key)

	var returnCharacter character.Character

	err = result.Scan(&returnCharacter.ID, &returnCharacter.Name, &returnCharacter.Class, &returnCharacter.Level, &returnCharacter.HitPointMaximum, &returnCharacter.Strength, &returnCharacter.Dexterity, &returnCharacter.Constitution, &returnCharacter.Wisdom, &returnCharacter.Intelligence, &returnCharacter.Charisma, &returnCharacter.Initiative, &returnCharacter.Modifiers)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(&returnCharacter.ID, &returnCharacter.Name, &returnCharacter.Level)

	json.NewEncoder(w).Encode(returnCharacter)

}
