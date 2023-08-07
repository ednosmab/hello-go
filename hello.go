package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	name := requestName()
	displayIntroduction(name)
	for {
		displayOptions(name)

		requestOption()
	}
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
	fmt.Println("Bem vindo ao Go World,", name)
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
	resp, _ := http.Get(site)
	return resp.StatusCode
}

func startMonitoring() {
	site := requestSite()
	if getStatusCode(site) != 200 {
		fmt.Println("Site verificado:", site, "Não foi carregado com sucesso")
	}
	fmt.Println("Site verificado:", site, "Carregado com sucesso!")
}
