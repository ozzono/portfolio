# Fundamentos

## Sumário

- [Constantes e Variáveis](#constantes-e-variáveis)
- [Tipos](#tipos)
  - [Tipos Básicos](#tipos-básicos)
    - [Array](#array)
    - [Slice](#slice)
    - [Struct](#struct)
    - [Map](#map)
    - [Channel](#channel)
    - [Function](#function)
    - [Methods](#methods)
    - [Interface](#interface)
  - [Os Zeros](#os-zeros)
- [Funções e Métodos](#funções-e-métodos)
  - [Sintaxes](#sintaxes)
- [Operadores Aritméticos](#operadores-aritméticos)
- [Operações de Atribuição](#operações-de-atribuição)
- [Operadores Relacionais](#operadores-relacionais)
- [Operadores Lógicos](#operadores-lógicos)
- [Operadores Unários](#operadores-unários)
- [Ponteiros](#ponteiros)

## Constantes e Variáveis

Golang, como a maioria das linguages, suporta variáveis e constantes. O que fazem é alto-explicativo: variáveis armazenam valores que podem ser modificados ao longo da execução do código enquanto constantes não.  
Constantes só podem ser de [tipo básico](#tipos-básicos).

##### [Sumário](#sumário)

## Tipos

Temos os tipos básicos e os tipos compostos da biblioteca padrão de Go.

##### [Sumário](#sumário)

## Tipos Básicos

Tipos básicos são aqueles que compõem os tipos nativos da linguagem.  
São eles:

- string: texto de qualquer tipo e sem tamanho definido;
- numéricos: entre os tipos numéricos temos subtipos:
  - inteiros: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64;
  - float (ponto flutuante): float32, float64;
  - complex (complexos): complex64, complex128;
- bool: booleano ou binário;
- byte: 8 bits com valores inteiros não negativos;
- rune: usado para caracteres e internamente são valores inteiros com 32 bits;

##### [Sumário](#sumário)

## Tipos compostos

### Array

Sequência de tamanho fixo com valores de mesmo tipo

- ex: novoArr:=[n]Tipo{v1,v2,v3,...}
  - n: tamanho do array;
  - Tipo: tipoe de valores na sequência;
  - v1, v2, v3: valores da sequência;

### Slice

Sequência de tamanho dinâmico que referencia os elementos de um array e tem declaração semelhante à deste.

- ex:

```go
    novoSlice:=[]Tipo{v1,v2,v3,..}
```

- Tipo: tipo de valores na sequência;
- v1, v2, v3: valores da sequência;

##### [Sumário](#sumário)

### Struct

Coleção de campos que podem ser de tipos diferentes.

- ex:

    ```go
        type Aluno struct {
            Nome string
            Idade int
            DataNascimento time.Time
        }
    ```

- Nome do tipo string
- Idade do tipo int
- DataNascimento do tipo time.Time (tipo padrão para datas)

##### [Sumário](#sumário)

### Map

Tipo similar a hash com referência do tipo chave/valor.

- ex:

    ```go
        aluno:=map[string]string{
            "nome":"José",
            "endereco":"Rua dos Bobos, n° 0",
        }
    ```

- nesse exemplo temos chaves do tipo string e valores também do tipo string.
- como acessar os valores:

    ```go
        nome:=aluno["nome"]         // José
        endereco:=aluno["endereco"] // Rua dos Bobos, n° 0
    ```

##### [Sumário](#sumário)

### Channel

Canais oferecem sincronização e comunicação entre gorotinas (goroutines). Considere como um túnel através do qual gorotinas enviam e recebem valores.  
O operador `<-` é usado para enviar ou receber dados.

- ex:

    ```go
        c <- v    // Enviado o valor de v para o canal c
        v := <- c // Recebendo o conteúdo do canal c e atribuindo à variável v
    ```

- Buffered and Unbuffer: Entre os tipos de canais temos os _buffered_ e _unbuffered_:
  - Unbuffered: Não tem informação de capacidade para armazenar valores, portanto envio e recebimento ficam bloqueados ou liberados a depender da rotina com a qual se comunica;
  - Buffered: Pode-se especificar o tamanho do buffer e usar o tamanho para realizar o controle a depender se o canal está cheio ou vazio;
- Pointer (ponteiros): Variável que armazena o endereço de memória de outra variável. O zero do ponteiro é `nil`
  - ex:

    ```go
        var texto="algun texto aqui"
        ponteiro :=*texto // armazena o valor do endereço em hexadecimal
    ```

##### [Sumário](#sumário)

### Function

Em go funções são valores e podem ser utilizadas como tal.

- ex:

    ```go
        func nomeDaFuncao(argumentos)retorno{
            // código da função
        }
    ```

##### [Sumário](#sumário)

### Methods

Possuem assinatura ligeiramente distinta das funções.  
Seguindo o exemplo abaixo, o recebedor pode ser uma struct ou uma interface ou qualquer outro tipo.  
O método tem acesso às propriedades do recebedor e pode chamar outros métodos.

- ex só para leitura:

    ```go
        func (recebedor tipoRecebedor) nomeDaFuncao(argumentos)retorno{
            // código da função
        }
    ```

- ex que permite escrita:

    ```go
        // ao declarar o método usando o ponteiro é possível modificar o recebor de forma que a modificação é consistente e acessível fora do método e após o seu encerramento
        func (recebedor *tipoRecebedor) nomeDaFuncao(argumentos)retorno{
            // código da função
        }
    ```

##### [Sumário](#sumário)

### Interface

É uma coleção de assinaturas de métodos. Todo tipo que implemente a assinatura da interface passa a ser também do tipo da interface.  
Seu valor zero é `nil`.

- ex:

    ```go
        type nomeDaInterface interface {
            Método1()                  // assinatura sem argumento e sem retorno
            Método2(argumentos)retorno // assinatura com argumento e retorno
        }
    ```

##### [Sumário](#sumário)

## Os Zeros

Em golang to tipo tem o chamado valor zero que é o valor atribuíndo à variável por padrão, ou seja, antes de se atribuir um valor ao longo da execução.  
Abaixo estão os zeros de cada tipo:

- numéricos:
  - int: 0 (para todos os tipos inteiros)
  - float: 0 (com as casas decimais, quando se aplica)
  - complex: 0 +0i
- bool: false (falso)
- byte: 0
- rune: 0

##### [Sumário](#sumário)

## Funções e Métodos

Funções e métodos compõem tipos, conforme mencionado [aqui](#function) e [aqui](#methods), e são parte fundamental para se escrever código em Go, especialmente quando se faz uso dos chamados _Design Patterns_, padrões de escrita de código que se propõem a resolver algum problema que faz parte do dia-a-dia com programação.  
Funções e métodos podem chamar uns aos outros com certa liberdade e para que sejam acessíveis externamente basta que sejam declarados com a primeira letra em maiúsculo.  
ex:

```go
    func NomeDaFuncao(){} // função exportada e acessível externamente, podendo ser usada bastando apenas importar o pacote que compõe
    func nomeDaFuncao(){} // função não exportada e inacessível externamente, podendo ser usada apenas pelo pacote que compõe
```

##### [Sumário](#sumário)

### Sintaxes

São as formas com que se pode escrever código de modo que o compilador compreenda e não acuse erros.
> Notas
>
> - É nas diferenças de sintaxe que se encontram o chamado `syntactic sugar` ou `açúcar sintático`. Denotam grau de conhecimento da linguagem e das diversas formas de se trabalhar bem como de escrever de forma idiomática.
> - É convenção na comunidade usar o chamado `camelCase` em que variáveis, funções e métodos são escritos sem espaço ou caractere para separar palavras, _ComoNesseExemplo_ e como em outros vistos nesse texto.

A seguir alguns exemplos:

- Declaração de variáveis - pode-se declarar variáveis das seguintes maneiras:

    ```go
        var variavel tipo // nesse caso é declarada a variável com tipo descrito e tem como valor os valores nulo que compõem o tipo
        variavel:=tipo{} // como no exemplo anterior é declarada a variável com os valores zero para cada tipo
        variavel:=funcaoComRetorno() // nesse caso a variável recebe o valor devolvido para funcão, sendo que o tipo é inferido automaticamente pelo compilador
        variavel:="algun texto aqui" // nesse caso a variável é declarar com valor e tipo definidos, sendo que o tipo é inferido pelo compilador
        var variavel = "algum texto aqui" // muito semelhante à declaração anterior, com a diferença de que o tipo não foi inferido, mas declarado
    ```

##### [Sumário](#sumário)

## Operadores Aritméticos

- `+` : realiza a soma de valores numéricos
- `-` : realiza a subtração de valores numéricos
- `*` : realiza a multiplicação de valores numéricos
- `/` : realiza a divisão de valores numéricos
- `%` : realiza a divisão entre valores e retorna o resto
  - ex:

    ```go
        x:= 10 % 2 // x igual a 2
        y:= 11 % 2 // x igual a 1
    ```

- `<<` : realiza a multiplicação do valor por `2` n vezes
  - ex:

    ```go
        x << n // multiplica x por 2 n vezes
    ```

- `>>` : realiza a divisão do valor por `2` n vezes
  - ex:

    ```go
        x >> n // divide x por 2 n vezes
    ```

##### [Sumário](#sumário)

## Operações de Atribuição

Operadores que atribuem valor a variáveis e constantes:

- `=`: atribui valor a variável já declarada;
- `:=`: atribui o valor e reserva memória, ou seja, é usado para declarar uma nova variável;
- `+=`: substitui valor da variável realizando a operação de `+`;
  - ex:

    ```go
    x:=10
    x+=2 // realiza em x a operação de +; x igual a 12
    ```

- `-=`: substitui valor da variável realizando a operação de `-`;
  - ex:

    ```go
    x:=10
    x-=2 // realiza em x a operação de -; x igual a 8
    ```

- `*=`: substitui valor da variável realizando a operação de `*`;
  - ex:

    ```go
    x:=10
    x*=2 // realiza em x a operação de *; x igual a 20
    ```

- `/=`: substitui valor da variável realizando a operação de `/`;
  - ex:

    ```go
    x:=10
    x/=2 // realiza em x a operação de /; x igual a 5
    ```

- `%=`: substitui valor da variável realizando a operação de `%`;
  - ex:

    ```go
    x:=10
    x%=2 // realiza em x a operação de %; x igual a 0
    ```

##### [Sumário](#sumário)

## Operadores Relacionais

Operadores comparativos entre valores.  
Tem por resultado um booleano.
> \* - operadores que só funcionam com valores numéricos

- `>`
  - maior que (*)
  - ex:

  ```go
  a > b // verdadeiro ou falso
  ```

- `>=`
  - maior ou igual (*)
  - ex:

  ```go
  a >= b // verdadeiro ou falso
  ```

- `<`
  - menor que (*)
  - ex:

  ```go
  a < b // verdadeiro ou falso
  ```

- `<=`
  - menor ou igual (*)
  - ex:

  ```go
  a <= b // verdadeiro ou falso
  ```

- `==`
  - igual
  - ex:

  ```go
  a == b // verdadeiro ou falso
  ```

- `!=`
  - diferente
  - ex:

  ```go
  a != b // verdadeiro ou falso
  ```

##### [Sumário](#sumário)

## Operadores Lógicos

Operadores booleanos, ou seja, operação com valores `verdadeiros` ou `falsos`.  
Cada operador bem como cada combinação de operadores tem sua respectiva tabela verdade, ou seja, considerado a combinação de todos os valores de entrada tem-se os valores de saída.  
Seguindo essa lógica, abaixo está a descrição dos operadores e de suas tabelas verdade (1 igual a verdadeiro e 0 igual a falso).  

---
Operador `||`
Em álgebra booleana esse operador é chamado de _OU_.
|OU|1|0|
|-|-|-|
|1|1|1|
|0|1|0|
ex:

```go
    a,b:=true,false
    x := a || b // valor de x é `true`
```

---
Operador `&&`
Em álgebra booleana esse operador é chamado de _E_.
|E|1|0|
|-|-|-|
|1|1|0|
|0|0|0|
ex:

```go
    a,b:=true,false
    x := a && b // valor de x é `false`
```

---

##### [Sumário](#sumário)

## Operadores Unários
Operadores que realizam incremento ou decremento de valores inteiros.
Note: Go não contém o operador ternário. A idéia é evitar que haja muitas formas de fazer a mesma operação básica a fim de tornar o código mais fácil de ler e manter.

- `++`: incrementa valores inteiros em 1 unidade
  - ex:

    ```go
        x:=1
        x++ // valor de x passa a ser 2
    ```

- `--`: decrementa valores inteiros em 1 unidade
  - ex:

    ```go
        x:=10
        x-- // valor de x passa a ser 9
    ```


##### [Sumário](#sumário)

## Ponteiros
