package main

import (
	"fmt"
	"go-orientacao-a-objetos/clientes"
	"go-orientacao-a-objetos/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoWillyan := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Willyan",
			CPF:       "150.574.041-11",
			Profissao: "DEV"},
		NumeroAgencia: 0001,
		NumeroConta:   155478}
	contaDoWillyan.Depositar(500)
	PagarBoleto(&contaDoWillyan, 250)

	fmt.Println(contaDoWillyan.Obtersaldo())

	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = clientes.Titular{Nome: "Silvia", CPF: "485.887.502-25", Profissao: "Engenheira"}
	contaDaSilvia.NumeroAgencia = 0001
	contaDaSilvia.NumeroConta = 441510205
	contaDaSilvia.Depositar(1000)
	PagarBoleto(&contaDaSilvia, 55)

	fmt.Println(contaDaSilvia.Sacar(400))
	fmt.Println(contaDaSilvia.Obtersaldo())
}
