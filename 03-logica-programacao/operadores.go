package main

import "fmt"

// baseado em: https://www.digitalocean.com/community/tutorials/understanding-boolean-logic-in-go-pt
func main() {
	// adicione um exemplo de chamada de cada metodo
	OpEquals(1, 1)
	OpNotEquals(1, 2)
	OpLess(1 , 2)
	OpGreater(2, 1)
	OpLessOrEqual(1, 2)
	OpGreaterOrEqual(2, 1)

	//
	fmt.Println("Posso fazer um saque?", PossoFazerUmSaque(100, 20))
	fmt.Println("Posso fazer um saque?", PossoFazerUmSaque(100, 100))
	fmt.Println("Posso fazer um saque?", PossoFazerUmSaque(100, 101))
	//fmt.Println("Posso fazer um saque?", StatusSaque(100, 20))
	fmt.Println("Posso fazer um saque?", StatusSaque(100, 100))
	//fmt.Println("Posso fazer um saque?", StatusSaque(100, 101))

}

func OpEquals(x int, y int) bool {
	fmt.Println("x == y:", x == y)
	return x == y
}

func OpNotEquals(x int, y int) bool {
	fmt.Println("x != y:", x != y)
	return x != y
}

func OpLess(x int, y int) bool {
	fmt.Println("x < y:", x < y)
	return x < y
}

func OpGreater(x int, y int) bool {
	fmt.Println("x > y:", x > y)
	return x > y
}

func OpLessOrEqual(x int, y int) bool {
	fmt.Println("x <= y:", x <= y)
	return x <= y
}

func OpGreaterOrEqual(x int, y int) bool {
	fmt.Println("x >= y:", x >= y)
	return x >= y
}

func PossoFazerUmSaque(saldo int, saque int) bool {
	return saldo >= saque
}

/*

Escreva nesta função os seguintes resultados para o calculo

	se o saldo for suficiente para o saque = SALDO_SUFICIENTE
	se o saldo for igual ao saque = SALDO_SUFICIENTE_IRA_ZERAR_O_SALDO
	se o saldo for menor que o saque = SEM_SALDO

*/
func StatusSaque(saldo int, saque int) string {

	if saldo == saque {
		return "SALDO_SUFICIENTE_IRA_ZERAR_O_SALDO"
	} else if saldo >= saque {
		return "SALDO_SUFICIENTE"
	} else {
		return "SEM_SALDO"
	}

}


