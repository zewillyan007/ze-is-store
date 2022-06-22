package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			imprimeLogs()
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
	sites := leSitesDoArquivo()
	for i := 0; i < 5; i++ {
		for i, site := range sites {
			fmt.Println()
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(10 * time.Second)
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site) //Função retorna duas variáveis, resp e err

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "tá a milhão!")
		registraLog(site, true)
		fmt.Println()
	} else {
		fmt.Println("Site:", site, "tá quebrado KKK! StatusCode:", resp.StatusCode)
		registraLog(site, false)
		fmt.Println()
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
