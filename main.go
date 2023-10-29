package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/wizsk/ihsan/data"
	"github.com/wizsk/ihsan/handlers"
)

func main() {
	// do stuff here
	path := "tmp.json"
	port := ":8001"

	db, err := data.OpenJDB(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.Index(w, r, db)
		}
	})

	http.HandleFunc("/api/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			_ = r.ParseForm()
			ar := r.FormValue("arabic")
			eng := r.FormValue("english")

			fmt.Println("/api/add", r.RemoteAddr)
			fmt.Println(ar, eng)
			db.Add(ar, eng)
		}
	})

	http.HandleFunc("/api/remove", nil)
	http.HandleFunc("/api/edit", nil)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}
