package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	CPF   int    `json:"cpf`
}

type ClienteInternacional struct {
	Cliente
	Pais string `json:"pais"`
}

func (c Cliente) exibe() {
	fmt.Println("Exibindo cliente pelo método", c.Nome)
}

func main() {
	cliente := Cliente{
		Nome:  "Robson",
		Email: "robsonfariasAD@gmail.com",
		CPF:   666666666655,
	}

	fmt.Println(cliente)

	cliente2 := Cliente{"Mari", "mari@gmail.com", 44343}

	fmt.Printf("Nome: %s Email: %s CPF: %d %s\n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Olavio",
			Email: "olavio@gmail.com",
			CPF:   4394933232943,
		},
		Pais: "USA",
	}

	fmt.Printf("Nome %s Email %s CPF %d %s\n", cliente3.Nome, cliente3.Email, cliente3.CPF)

	cliente.exibe()
	cliente2.exibe()
	cliente3.exibe()

	clienteJson, err := json.Marshal(cliente3)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(clienteJson))

	jsonCliente4 := `{"nome":"Olavio","email":"olavio@gmail.com","CPF":4394933232943,"pais":"USA"}`
	cliente4 := ClienteInternacional{}

	json.Unmarshal([]byte(jsonCliente4), &cliente4)

	fmt.Println(cliente4)
}
