package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Pessoa struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	http.HandleFunc("/greetings", func(w http.ResponseWriter, r *http.Request) {
		var p Pessoa
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&p)
		if err != nil {
			http.Error(w, "Erro ao decodificar o JSON", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello " + p.FirstName + " " + p.LastName))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
