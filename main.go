package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/smwest87/dnd-golang/api"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/home", api.HomeLink)
	router.HandleFunc("/character/{id}", api.GetCharacter).Methods("GET")
	router.HandleFunc("/character/{id}", api.DeleteCharacter).Methods("DELETE")
	router.HandleFunc("/character/[{id}", api.UpdateCharacter).Methods("PUT")
	router.HandleFunc("/character/new", api.CreateCharacter).Methods("POST")

	/*hero, err := character.GenerateCharacter(os.Args[1], os.Args[2])
	_, err = character.InsertCharacter(*hero)
	if err != nil {
		log.Fatal(err)
	} */

	fmt.Println("Preparing to serve")
	http.ListenAndServe(":10000", router)
	fmt.Println("serve failed")

}
