package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Id int `json:"id"`
	PersonRecieve
}

type PersonRecieve struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
}

type ServicePerson map[int]Person

type ServicePersonInterface interface {
	Create(p Person) error
	Get(w http.ResponseWriter)
	List(id int, w http.ResponseWriter)
	Update(id int, w http.ResponseWriter)
	delete(id int, w http.ResponseWriter)
}

func (ps ServicePerson) Create(recieve PersonRecieve) error {

	if err := recieve.Validate(); err != nil {
		return fmt.Errorf("impossível criar o usuário %w", err)
	}

	id := len(ps) + 1
	ps[id] = Person{
		Id:            id,
		PersonRecieve: recieve,
	}

	return nil

}

func (ps ServicePerson) List(w http.ResponseWriter) {

	if len(ps) > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ps)
		return
	}

}

func (ps ServicePerson) Get(id int, w http.ResponseWriter) {

	p, prs := ps[id]
	if prs {
		json.NewEncoder(w).Encode(&p)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Message{
		Message: "Pessoa não encontrada, por favor verifique os campos",
	})

}

func (ps ServicePerson) Update(id int, p PersonRecieve, w http.ResponseWriter) {

	existing, prs := ps[id]
	if prs {

		oldName := existing.Name
		if len(p.Name) > 0 {
			existing.Name = p.Name
		}
		if len(p.Email) > 0 {
			if err := validEmail(p.Email); err != nil {
				json.NewEncoder(w).Encode(Message{
					Message: "Não foi possível atualizar o usuário",
					Error:   err.Error(),
				})
			}

			existing.Email = p.Email
		}
		if len(p.Cpf) > 0 {

			if err := validCPF(p.Cpf); err != nil {
				json.NewEncoder(w).Encode(Message{
					Message: "Não foi possível atualizar o usuário",
					Error:   err.Error(),
				})
			}

			existing.Cpf = p.Cpf
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Message{
			Message: "Usuario " + oldName + " atualizado com suceso",
		})
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(Message{
		Message: "Usuario não existe!",
	})

}

func (ps ServicePerson) Delete(id int, w http.ResponseWriter) {

	_, prs := ps[id]
	if prs {
		delete(ps, id)
		for i := id + 1; i < len(ps); i++ {
			ps[id] = ps[id-1]
		}
		return
	}

	json.NewEncoder(w).Encode(Message{
		Message: "Essa pessoa não está registrada",
		Error:   "ID não encontrado",
	})
	return

}
