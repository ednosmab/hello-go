package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitorings = 5
const delay = 3
const delayMessage = 1

func main() {
	leArquivoExterno()
	// name := requestName()
	// displayIntroduction(name)
	// for {
	// 	displayOptions(name)

	// 	requestOption()
	// }
}

func informName() {
	fmt.Println("Informe o seu nome")
}
func requestName() string {
	var name string
	informName()
	fmt.Scan(&name)
	return name
}
func displayIntroduction(name string) {
	version := 1.1
	fmt.Println("Bem vindo ao Site Monitor,", name)
	fmt.Println("Este programa está na versão", version)
}

func displayOptions(name string) {
	fmt.Println(name, "informe a opção desejada:")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Log")
	fmt.Println("0- Sair do Programa")
}

func informOption() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func requestOption() {
	command := informOption()
	switch command {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando!")
		os.Exit(-1)
	}
}

func requestSite() string {
	fmt.Println("Informe o site a ser verificado:")
	var site string
	fmt.Scan(&site)
	return site
}
func getStatusCode(site string) int {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return resp.StatusCode
}

func allSites() []string {
	sites := []string{
		"https://www.google.com.br",
		"https://www.casadocodigo.com.br/",
		"http://alura.com.br/",
	}

	return sites
}

func siteValidate(site string) {
	if getStatusCode(site) != 200 {
		fmt.Println("Site verificado:", site, "Não foi carregado com sucesso")
	}
	fmt.Println("Site verificado:", site)
	fmt.Println("Status: carregado com sucesso!")
}
func processSite() {
	sites := allSites()
	for i := 0; i < monitorings; i++ {
		for _, site := range sites {
			siteValidate(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
	fmt.Println("-----------Menu de opções------------")
}
func startMessage() {
	fmt.Println("")
	fmt.Print("Monitorando ")
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(delayMessage * time.Second)
	}
}
func startMonitoring() {
	startMessage()
	processSite()
}

func leArquivoExterno() {
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		return
	}
	fmt.Println(arquivo)
}
