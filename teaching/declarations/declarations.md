# Breve demonstração com formas de declarar variáveis

## Go suporta algumas formas distintas de declarar variáveis ilustradas abaixo

[declarations.go](./declarations.go)
```go
 // nessas duas formas a variável é declarada com os valores nulos para os tipos primitivos
 var variavelMagica tipoMagico

 // nessas duas formas a variável é declarada com os valores nulos para os tipos primitivos
 outraMagia := tipoMagico{}

 maisUmaMagia := tipoMagico{
  campoMagico:    "haja magia",
  medidorDeMagia: 8000,
 }

 retornoMagico, erroMagico := funcaoMagica()

 // declaração anônima não funciona sem declarar valores, senão seria a declaração de um tipo
 // o tipo não pode ser anônimo, no máximo pode ser interno ao escopo
 // antepenultimaMagia:=struct type{
 // campoMagico string
 // medidorDeMagia int
 // }

 penultimaMagia := struct {
  campoMagico    string
  medidorDeMagia int
 }{
  campoMagico:    "já deu, né?",
  medidorDeMagia: 0,
 }

 ultimaMagia := struct {
  campoMagico    string
  medidorDeMagia int
 }{
  "#shrug",
  0,
 }

 var pareiComAMagia = tipoMagico{
  "deixa pra lá",
  0,
 }

```

Retorno no terminal

```log
variavelMagica:  { 0}
variavelMagica: main.tipoMagico{campoMagico:"", medidorDeMagia:0}
outraMagia:  { 0}
outraMagia: main.tipoMagico{campoMagico:"", medidorDeMagia:0}
maisUmaMagia:  {haja magia 8000}
maisUmaMagia: main.tipoMagico{campoMagico:"haja magia", medidorDeMagia:8000}
retornoMagico:  { 0}
retornoMagico: main.tipoMagico{campoMagico:"", medidorDeMagia:0}
erroMagico:  <nil>
penultimaMagia:  {já deu, né? 0}
penultimaMagia: struct { campoMagico string; medidorDeMagia int }{campoMagico:"já deu, né?", medidorDeMagia:0}
ultimaMagia:  {#shrug 0}
ultimaMagia: struct { campoMagico string; medidorDeMagia int }{campoMagico:"#shrug", medidorDeMagia:0}
pareiComAMagia:  {deixa pra lá 0}
pareiComAMagia: main.tipoMagico{campoMagico:"deixa pra lá", medidorDeMagia:0}
```
