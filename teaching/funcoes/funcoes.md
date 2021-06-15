# Breve demonstração de formas de declarar funcoes

[funcoes.go](./funcoes.go)

```go

type funcional struct{ string }

func funcoes() {
 fmt.Println("essa função faz muita coisa")
 outraFuncao := func() {
  fmt.Println("essa função também faz muita coisa")
 }
 outraFuncao()
 f := funcional{"teste"}
 f.Imprimir("argumento espetacular")
}

func (f *funcional) Imprimir(arg string) {
 fmt.Printf("argumento: %s", arg)
 fmt.Printf("funciona!: %s", f)
}
```

saída:

```log
essa função faz muita coisa
essa função também faz muita coisa
argumento: argumento espetacular
funciona!: teste
PASS
ok  	teaching/funcoes	0.001s
```
