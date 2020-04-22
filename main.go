package main

import (
	"fmt"
	"log"
	"net/http"

	dbase "./dbase"
	mplx "./mplx"
)

func main() {
	db, err := dbase.Create("forumDB")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		mplx.Multiplexer(db, w, r)
	})

	fmt.Println("Server is listening to port :8080...")
	http.ListenAndServe(":8080", nil)
}
