package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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

func messageMenu() {
	fmt.Println("-----------Menu de opções------------")
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
		displayLogs()
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
		logRecord(site, false)

	} else {
		fmt.Println("Site verificado:", site)
		fmt.Println("Status: carregado com sucesso!")
		logRecord(site, true)
	}

}

func processSite() {
	sites := readExternalFile("sites.txt")
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

	messageMenu()
}

func startMessage(option string) {
	fmt.Println("")
	fmt.Print(option, " ")

	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(delayMessage * time.Second)
	}
}

func readExternalFile(fileToRead string) []string {
	lines := []string{}
	file, err := os.Open(fileToRead)
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

func logRecord(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString("[ " + time.Now().Format("02/01/2006 15:04:05") + " ] -- " + site + " -- online: " + strconv.FormatBool(status) + "\n")

	file.Close()

}

func startMonitoring() {
	startMessage("Monitorando")
	processSite()
}

func readExternalLog(fileToRead string) {
	file, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))

	messageMenu()
}
func displayLogs() {
	startMessage("Exibindo logs")
	readExternalLog("log.txt")

}
