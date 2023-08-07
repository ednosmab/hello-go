package main

import (
	"fmt"
	"os"
)

func main() {
	nome := solicitaNome()
	exibeIntroducao(nome)
	exibeOpcoes(nome)

	solicitaOpcao()
}

func informeNome() {
	fmt.Println("Informe o seu nome")
}
func solicitaNome() string {
	var nome string
	informeNome()
	fmt.Scan(&nome)
	return nome
}
func exibeIntroducao(nome string) {
	versao := 1.1
	fmt.Println("Bem vindo ao Go World,", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeOpcoes(nome string) {
	fmt.Println(nome, "informe a opção desejada:")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Log")
	fmt.Println("0- Sair do Programa")
}

func informeOpcao() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func solicitaOpcao() {
	comando := informeOpcao()
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
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
