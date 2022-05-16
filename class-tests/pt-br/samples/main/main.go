package main

// pacote principal da linguagem go
// por onde todo código Go começa

import "fmt"

// entre os imports incluímos pacotes que serão utilizados pelo código
// nesse exemplo importamos apenas o fmt, pacote utilizado para formatar texto e imprimir no terminal

// função principal do código em Go por onde tudo começa
// - só pode have uma função `main` no pacote `main`
// - para que o código seja executado é preciso que haja um pacote main e uma função main
func main() {
	fmt.Println("Olá Gopher")
	// Gopher é tanto o nome do mascote do Go,
	// também uma tradição em tecnologias de código aberto,
	// e também é o apelido atribuído a desenvolvedores que usam a linguagem
}
