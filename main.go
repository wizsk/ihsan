package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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
		if r.URL.Path != "/" {
			http.Error(w, "bad req", 400)
			return
		}
		if r.Method == http.MethodGet {
			handlers.Index(w, r, db)
		}
	})

	http.HandleFunc("/api/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// err ignoring
			_ = r.ParseForm()
			ar := r.FormValue("arabic")
			eng := r.FormValue("english")
			harakat := r.FormValue("respect_harakats")

			var harakats bool
			if harakat == "true" {
				harakats = true
			} else if harakat == "false" {
				harakats = false
			} else {
				http.Error(w, "invalid form value", http.StatusBadRequest)
				return
			}

			fmt.Println("/api/add", r.RemoteAddr, ar, eng, harakat)
			if err := db.Add(ar, eng, harakats); err != nil {
				if v, ok := err.(*data.Vocab); ok {
					http.Error(w, fmt.Sprintf(`{"err": "%v is already in the database", "data": {"id": %d, "arabic": "%s", "english": "%s"}}`, v, v.Id, v.Arabic, v.English), http.StatusBadRequest)
					return
				}

				http.Error(w, `{"err": "unknown error"}`, http.StatusBadRequest)
				fmt.Println(err)
				return
			}

			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}
	})

	http.HandleFunc("/api/remove", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}

		r.ParseForm()
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println(err)
			return
		}

		if err := db.Remove(id); err != nil {
			log.Println(err)
			return
		}
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	})

	http.HandleFunc("/api/edit", func(w http.ResponseWriter, r *http.Request) {
		// fields
		// id, arabic, english
		// look for duplicates // perhaps

		_ = r.ParseForm()

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println(err)
			return
		}
		arabic := r.FormValue("arabic")
		englsih := r.FormValue("english")

		fmt.Println("/api/add", r.RemoteAddr, id, arabic, englsih)
		if err := db.Edit(id, arabic, englsih); err != nil {
			log.Println(err)
			return
		}
	})

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}
