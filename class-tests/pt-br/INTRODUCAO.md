# Introdução

## Sumário

- [O que é Go](#o-que-é-go-golang)
- [Olá Gopher](#olá-gopher)

## O que é Go (Golang)

Numa breve descrição, Golang é uma linguagem inicialmente desenvolvida pelo Google e posteriormente tornada de código aberto.  
É compilada e tem como propósito ser uma linguagem leve, rápida e simples de ler e escrever.

## Olá gopher

Seguindo a tradição da maioria dos tutoriais, o primeiro código escrito compreende a escrita de um cumprimento. Abaixo vemos como é escrito usango Go.
> [Código de exemplo](/samples/main/main.go)

```go
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
```

Retorno do comando `./samples/main/main` ou `make main`

```log
$ make main                     
./samples/main/class-test-main
Olá Gopher
```
