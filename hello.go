package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoring = 5
const delay = 3
const delayMessage = 1

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
	site = "https://" + site
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	return resp.StatusCode
}

func siteValidate(site string) {
	if getStatusCode(site) != 200 {
		fmt.Println("Site verificado:", site, "Não foi carregado com sucesso")
	} else {
		fmt.Println("Site verificado:", site)
		fmt.Println("Status: carregado com sucesso!")
	}

}

func processSite() {
	sites := readExternalFile()
	fmt.Println(sites)
	for i := 0; i < monitoring; i++ {
		if len(sites) == 0 {
			break
		}
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

func readExternalFile() []string {
	lines := []string{}
	file, err := os.Open("sites.txt")
	reader := bufio.NewReader(file)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		return lines
	}

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		lines = append(lines, line)

		if err == io.EOF {
			break
		}

	}

	file.Close()
	return lines
}

func startMonitoring() {
	startMessage()
	processSite()
}
