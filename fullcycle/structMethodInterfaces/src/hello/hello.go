package main

import "fmt"

type Veiculo interface {
	buzina()
}

func (c Carro) buzina() {
	fmt.Println(c.Fabricante, "buzinou")
}

func (m Moto) buzina() {
	fmt.Println(m.Fabricante, "buzinou")
}

type Pessoa struct {
	Nome    string
	Idade   int
	Veiculo Veiculo
}
type Moto struct {
	Fabricante string
	Ano        int
}

type Carro struct {
	Fabricante string
	Modelo     string
	Ano        int
}

func (p Pessoa) andou() {
	p.Nome = "viado"
	fmt.Println(p.Nome, "andou")
}

func main() {
	// bmw := Carro{
	// 	Fabricante: "bmw",
	// 	Modelo:     "xpto",
	// 	Ano:        2021,
	// }

	moto := Moto{
		Fabricante: "moto bmw",
		Ano:        2021,
	}

	robson := Pessoa{
		Nome:    "Robson",
		Idade:   23,
		Veiculo: moto,
	}

	robson.andou()
	robson.Veiculo.buzina()
	fmt.Println("Nome:", robson.Nome)
}
