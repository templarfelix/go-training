package main

import (
	"fmt"
	"log"
	"time"
)

type client struct {
	name string // fernando luís
	documentNumber string // 053.043.099-xx
	clientType clientType //
	gender string // male
	address address //
	email string // felix.templar@gmail.com
	phoneNumber string // +55 47 9-9919-8914
	account []account
}

type clientType struct {
	legalPerson bool // false
}

type address struct {
	address string // rua porto união
	number string // 1237
	complement string // ap 804
	cep string // 89203-460
	neighborhood string // anita
	city string // joinville
	state string // santa catarina
	country string // brasil
}

type account struct {
	accountNumber string // 033-01-01
	accountType string // salario
	saldo float64 // 0
	limit float64 // 1000

	// representa o objeto transferencia
	transfer []transfer
	//
	extrato []extrato
}

type transfer struct {
	transferType string // TED || DOC || PIX || ...
	time time.Time
	accountOrigin account // fernando
	accountDestino account // titi
	value float64
	transferIndex float64
}

type extrato struct {
	saldoDoDia float64

	time time.Time
	tipoSaida string
	tipoEntrada string
	movimentacoes []movimentacoes

}

type movimentacoes struct {

}

//
type microondas struct {
	marca string
	voltagem string
	capacidade string
	cor string
}

type pessoa struct {
	nome string
	genero string
	peso float64
	altura float64
	dataNascimento time.Time
	nascionalidade string // brasileiro
	naturalidade string // joinville
	numeroDocumento string
}

func main() {

	location, err:= time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatalf("ERRRO NA GERACAO DO LOCATION")
	}
	alvaro := pessoa{nome: "Alvaro", genero: "masculino", peso: 70.00, altura: 1.75,
			dataNascimento: time.Date(1991, time.February, 22, 01, 01, 01, 00, location),
			nascionalidade: "peruano", naturalidade: "capichaba", numeroDocumento: "xxx-xxx",
	}

	fmt.Println(alvaro.nome)
	fmt.Println(alvaro.dataNascimento)
	fmt.Println("Dia" , alvaro.dataNascimento.Day())
	fmt.Println(alvaro.dataNascimento.Year())
	fmt.Println("a idade da pessoal é ", calculaIdadeDaPessoa(alvaro))
	fmt.Println(alvaro.nome)
}

func calculaIdadeDaPessoa(pessoa pessoa) int  {
	return 2021 - pessoa.dataNascimento.Year()
}