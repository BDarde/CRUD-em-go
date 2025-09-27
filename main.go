package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Message struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func main() {

	var service = make(ServicePerson)

	handle := func(w http.ResponseWriter, r *http.Request) {
		var p PersonRecieve
		var responseEncoder = json.NewEncoder(w)
		var bodyDecod = json.NewDecoder(r.Body)

		switch r.Method {
		case http.MethodPost:

			//validação content-type
			if r.Header.Get("Content-Type") != "application/json" {
				responseEncoder.Encode(Message{Error: "esperava um json"})
				return
			}

			if err := bodyDecod.Decode(&p); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				responseEncoder.Encode(Message{
					Message: "Falhou ao decodificar os valores",
					Error:   err.Error(),
				})
				return
			}

			//validção campos
			if err := p.Validate(); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				responseEncoder.Encode(Message{
					Error: err.Error(),
				})
				return
			}

			if err := service.Create(p); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				responseEncoder.Encode(Message{
					Message: "Usuário criado com sucesso",
				})
				return
			}

			w.WriteHeader(http.StatusCreated)
			responseEncoder.Encode(Message{
				Message: "Usuário criado com sucesso",
			})
			return

		case http.MethodGet:

			url := r.URL
			if id := url.Query().Get("id"); id != "" {

				id, _ := strconv.Atoi(id)
				service.Get(id, w)
				return
			}

			service.List(w)
			return

		case http.MethodPut:

			url := r.URL
			if id := url.Query().Get("id"); id != "" {

				//validação content-type
				if r.Header.Get("Content-Type") != "application/json" {
					responseEncoder.Encode(Message{Error: "esperava um json"})
					return
				}

				if err := bodyDecod.Decode(&p); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					responseEncoder.Encode(Message{
						Message: "Falhou ao decodificar os valores",
						Error:   err.Error(),
					})
					return
				}

				id, _ := strconv.Atoi(id)
				service.Update(id, p, w)
				return
			}

		}

	}

	http.HandleFunc("/person", handle)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
