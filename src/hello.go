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

var monitoramentos int = 3

const delay = 5

func main() {

	exibeIntro()
	for {
		exibeMenu()
		comando := leComando()
		// *************************************** PONTOS DE ATENÇÃO ***************************************
		// Valores não atribuidos a linguagem começa com o valor zerado, 0, 0.0, caracter vazio
		// Todas as variaveis declaradas devem ser usadas
		// Todas as funções de pacotes externos devem começar com Maiusculo
		// Variaveis conseguem descobrir o tipo sozinho, enferencia dos tipos das variaveis
		// Go possui operador de atribuição de variaveis = 	nome := "Thiago"
		//													versao := 1.1
		//													idade := 22
		// Não tem while, loops é com for
		// arrays tem tamanhos fixos, slices são os substitutos com tamanho indefinido
		// *************************************************************************************************

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 3:
			imprimeSites()
		case 4:
			alteraQtd()
		case 9:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não reconheco esse comando")
			os.Exit(-1)
		}
	}
}

func exibeIntro() {
	var nome string = "Thiago"
	var versao float32 = 1.2

	fmt.Println("")
	fmt.Println("Olá,", nome)
	fmt.Println("Programa está na versão:", versao)
	fmt.Println("")

}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Exibir Sites")
	fmt.Println("4 - Alterar qtd de testes")
	fmt.Println("9 - Sair do programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("")
	return comando
}

func iniciarMonitoramento() {

	fmt.Println("Monitorando...")

	//sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br/", "https://www.caelum.com.br"}
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu o erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site,", site, "foi carregado com sucesso.")
		registraLog(site, true)
	} else {
		fmt.Println("O site", site, "esta com problemas")
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu o erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n') //retorna uma string para cada linha
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func alteraQtd() {
	fmt.Println("Digite a quantidade de vezes que deseja testar")
	fmt.Scan(&monitoramentos)
	fmt.Println("")
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.csv", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu o erro: ", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + ";" + site + ";" + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.csv")

	if err != nil {
		fmt.Println("Ocorreu o erro: ", err)
	}
	fmt.Println(string(arquivo))
	fmt.Println("")
}
func imprimeSites() {

	arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu o erro: ", err)
	}
	fmt.Println(string(arquivo))
	fmt.Println("")
}
