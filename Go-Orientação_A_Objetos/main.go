package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoWillyan := ContaCorrente{titular: "Willyan",
		numeroAgencia: 0001, numeroConta: 155478, saldo: 135.5}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println(contaDoWillyan)
	fmt.Println(contaDaBruna)
}
