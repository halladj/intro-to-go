package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"web_server/util"
)

type Text struct {
	Text string
	Key  string
}

var u util.Util = util.NewUtil()

func EncryptionHandler(w http.ResponseWriter, r *http.Request) {
	var t Text

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "CipherText>> "+u.VigenereEncryption(t.Text, t.Key))

}

// TODO Define the DecryptionHandler
//

func main() {
	var m *http.ServeMux = http.NewServeMux()

	m.HandleFunc("/encrypt", EncryptionHandler)
	// TODO Add new route called << decrypt>>
	//      and associate it with DecryptionHandler

	s := &http.Server{
		Addr:    ":8080",
		Handler: m,
	}

	log.Fatal("Server Listenning on Port: 8080", s.ListenAndServe())
}
