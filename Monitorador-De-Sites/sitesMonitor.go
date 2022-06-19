package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	intro()

	for {
		menu()

		comando := controll()

		switch comando {
		case 1:
			monitoramento()
		case 2:
			fmt.Println("Exibindo Log...")
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Cê é burro cara que loucura.")
			os.Exit(-1)
		}
	}
}

func intro() {
	var nome string
	versao := 1.1
	fmt.Print("Qual o seu nome? ")
	fmt.Scan(&nome)
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func menu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println()
	fmt.Println("0 - Sair do Programa")
}

func controll() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	return comando
}

func monitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://alura.com.br", "https://youtube.com.br"}

	for i := 0; i < 5; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(10 * time.Second)
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site) //Função retorna duas variáveis, resp e err
	//fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "tá a milhão!")
		fmt.Println()
	} else {
		fmt.Println("Site:", site, "tá quebrado KKK! StatusCode:", resp.StatusCode)
		fmt.Println("")
	}
}
