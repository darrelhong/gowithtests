package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := os.OpenFile("game.db.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening game.db.json %v", err)
	}

	store := NewFileSystemPlayerStore(db)

	server := NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))

}
