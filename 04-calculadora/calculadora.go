package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println("soma", soma(1, 1))
	//fmt.Println("subtracao", subtracao(1, 1))
	//fmt.Println("subtracao", subtracao(1.1, 1))
	//fmt.Println("adicionarPercentual", adicionarPercentual(20, 100))
	//fmt.Println("descobrirPercentual", descobrirPercentual(100, 10))

	//var teste = soma(subtracao(10, multiplicacao(10 , 10)) , 20)
	//fmt.Println(teste)
	//multiplicacao(10, 10)
	var result float64 = soma(10, 10.2)
	if result == 20.2 {
		fmt.Println("ebaaa")
	} else {
		fmt.Println("fuuuu")
	}
}

func soma(x float64, y float64) float64 {
	return x + y
}

func subtracao(x float64, y float64) float64 {
	return x - y
}

func multiplicacao(x float64, y float64) float64 {
	return x * y
}

func divisao(x float64, y float64) float64 {
	return x / y
}

func adicionarPercentual(valor float64, porcentagem float64) float64 {
	return valor + (valor * porcentagem / 100)
}

func descobrirPercentual(valor float64, valor2 float64) float64 {
	return (valor2 / valor) * 100
}

func pow(valor float64, multiplo float64) float64 {
	return math.Pow(valor, multiplo)
}
