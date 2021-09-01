package main

import "fmt"

// FIXME criar o objeto Pessoa
// irmaos
// pais
// pets

// criar o metodo criaEU
//

type pessoa struct {
	nome string // fernando luís
	idade int
	filho []pessoa
	irmao []irmao
	pais []pessoa
	pet []animal
}

type irmao struct {
	nome string
	idade int
	filho []pessoa
}

type animal struct {
	nome  string
	raca  string
	idade int
}

func main() {

	fernando := pessoa{nome: "Fernando Luís da silva", idade: 34,
		pet: []animal{{nome: "luna", raca: "canina", idade: 5}},
		irmao: []irmao{{nome: "Fabiane", idade: 39,
								filho: []pessoa{{nome: "Murilo", idade: 9},
									            {nome: "Raul", idade: 6}}},
						{nome: "Fellipe", idade: 32}},
		pais: []pessoa{{nome: "Jose da Silva", idade: 61, pet: nil, irmao: nil},
					   {nome: "Dervania Doroti Drase da Silva", idade: 55}},
	}
	fmt.Println(fernando)

}