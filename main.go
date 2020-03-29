package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lukehinds/godirectory/db"
)

// https://siongui.github.io/2016/01/09/go-sqlite-example-basic-usage/

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

type Listing struct {
	ID          int16  `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

func createListing(w http.ResponseWriter, r *http.Request) {
	var newListing Listing
	json.NewDecoder(r.Body).Decode(&newListing)
	fmt.Fprintf(w, "ID: %v\n", newListing.ID)
	fmt.Fprintf(w, "Name: %s\n", newListing.Name)
	fmt.Fprintf(w, "Desc: %s\n", newListing.Description)
}

func main() {
	const dbpath = "foo.db"
	db := db.InitDB(dbpath)
	defer db.Close()
	CreateTable(db)
	router := chi.NewRouter()
	router.Get("/", homeLink)
	router.Post("/listing", createListing)
	// router.Route("/listings", func(router chi.Router) {
	//	router.Get("/", getAllListings)
	//	router.Get("/{id}", getOneListing)
	//	router.Patch("/{id}", updateListing)
	//	router.Delete("/{id}", deleteListing)
	// })

	log.Fatal(http.ListenAndServe(":8080", router))
}
