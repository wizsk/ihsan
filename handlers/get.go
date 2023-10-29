package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/wizsk/ihsan/data"
)

func Index(w http.ResponseWriter, r *http.Request, db *data.JDB) {
	fmt.Println("index", r.RemoteAddr)
	buf := new(bytes.Buffer)
	for _, v := range db.GetVocabs().Words {
		fmt.Fprintf(buf, "arabic: %s\tenglish: %s\n", v.Arabic, v.English)
	}
	w.Write(buf.Bytes())
}
