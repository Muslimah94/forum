package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendJSON function marshals and sends given data to response writer
func SendJSON(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	//err := json.NewEncoder(w).Encode(v)
	if err != nil {
		fmt.Println("SendJSON json.Marshal ERROR:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// ReceiveJSON function decodes data from request
func ReceiveJSON(r *http.Request, v interface{}) {
	var w http.ResponseWriter
	err1 := json.NewDecoder(r.Body).Decode(v)
	if err1 != nil {
		fmt.Println("ReceiveJSON: Failed to Decode", err1)
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
}
