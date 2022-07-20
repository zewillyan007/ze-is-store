package main

import (
	"fmt"
	"go-orientacao-a-objetos/contas"
)

func main() {
	contaDoWillyan := contas.ContaCorrente{Titular: "Willyan",
		NumeroAgencia: 0001, NumeroConta: 155478, Saldo: 135.5}

	contaDaBruna := contas.ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println(contaDoWillyan)
	fmt.Println(contaDaBruna)

	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 500

	fmt.Println(contaDaSilvia.Saldo)

	fmt.Println(contaDaSilvia.Sacar(400))
	fmt.Println(contaDaSilvia.Saldo)
}
