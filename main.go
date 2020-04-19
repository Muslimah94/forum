package main

import (
	"fmt"
	"net/http"

	mplx "./mplx"
)

func main() {
	http.HandleFunc("/api", mplx.Multiplexer)

	fmt.Println("server is listening to port :8080....")
	http.ListenAndServe(":8080", nil)
}
