package handlers

import (
	"log"
	"net/http"
)

func errToHttp(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), 500)
	log.Println(err)
}
