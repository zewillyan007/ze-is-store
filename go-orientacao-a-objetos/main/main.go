package main

import (
	"fmt"
	"go-orientacao-a-objetos/clientes"
	"go-orientacao-a-objetos/contas"
)

func main() {
	contaDoWillyan := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Willyan",
			CPF:       "150.574.041-11",
			Profissao: "DEV"},
		NumeroAgencia: 0001,
		NumeroConta:   155478}

	fmt.Println(contaDoWillyan.Obtersaldo())

	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = clientes.Titular{Nome: "Silvia", CPF: "485.887.502-25", Profissao: "Engenheira"}
	contaDaSilvia.NumeroAgencia = 0001
	contaDaSilvia.NumeroConta = 441510205

	fmt.Println(contaDaSilvia.Sacar(400))
	fmt.Println(contaDaSilvia.Obtersaldo())
}
