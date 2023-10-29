package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/wizsk/ihsan/data"
)

func Index(w http.ResponseWriter, r *http.Request, db *data.JDB) {
	tmpl := template.New("index").Funcs(template.FuncMap{
		"formatTime": func(t time.Time) string {
			return t.Format("03:04 PM Mon Jan")
		},
	})

	_, err := tmpl.ParseFiles("static/frontend/index.html", "static/frontend/index.js")
	if err != nil {
		panic(err)
	}

	fmt.Println("index", r.RemoteAddr)
	buf := new(bytes.Buffer)
	for _, v := range db.GetVocabs().Words {
		fmt.Fprintf(buf, "arabic: %s\tenglish: %s\n", v.Arabic, v.English)
	}
	if err := tmpl.ExecuteTemplate(w, "index.html", db.GetVocabs().Words); err != nil {
		panic(err)
	}

	buf.WriteString(`<form action="/api/add" method="post">
        <label for="arabic">Arabic Word:</label>
        <input type="text" id="arabic" name="arabic" required>
        <br><br>

        <label for="english">English Translation:</label>
        <input type="text" id="english" name="english" required>
        <br><br>

        <input type="submit" value="Submit">
    </form>
`)
	// w.Write(buf.Bytes())
}
