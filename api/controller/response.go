package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteSuccess(w http.ResponseWriter, data interface{}) {
	bytes,  err := json.Marshal(data)
	if err != nil {
		log.Fatal("don't parse bytes")
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	//w.Write(structs)
}
