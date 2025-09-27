package main

import (
	"errors"
	"fmt"
	"strings"
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

func (p PersonRecieve) Validate() error {

	if len(p.Name) == 0 {
		return fmt.Errorf("nome não pode vir vazio")
	}

	if err := validCPF(p.Cpf); err != nil {
		return err
	}

	if err := validEmail(p.Email); err != nil {
		return err
	}

	return nil
}

func validEmail(str string) error {

	err := errors.New("não é um email válido")
	if !strings.Contains(str, "@") {
		return err
	} else {
		if !strings.Contains(strings.Split(str, "@")[1], ".com") {
			return err
		}
	}

	return nil

}

func validCPF(str string) error {
	if len(str) == 0 {
		return fmt.Errorf("cpf não pode ser vazio")
	}

	// caso venha com caracteres especiais ".", "-"
	if len(str) < 11 || len(str) > 14 {
		return fmt.Errorf("cpf está no formado incorreto")
	}

	return nil

}
