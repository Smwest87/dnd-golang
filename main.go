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
	router.HandleFunc("/character/{id}", api.ResponseWrapper(api.GetCharacter)).Methods("GET")
	router.HandleFunc("/character/{id}", api.ResponseWrapper(api.DeleteCharacter)).Methods("DELETE")
	router.HandleFunc("/character/[{id}", api.ResponseWrapper(api.UpdateCharacter)).Methods("PUT")
	router.HandleFunc("/character/new", api.ResponseWrapper(api.CreateCharacter)).Methods("POST")
	fmt.Println("Preparing to serve")
	http.ListenAndServe(":10000", router)
	fmt.Println("serve failed")
}
