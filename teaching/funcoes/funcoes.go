package funcoes

import "fmt"

type funcional struct{ string }

func funcoes() error {
	fmt.Println("essa função faz muita coisa")

	outraFuncao := func() {
		fmt.Println("essa função também faz muita coisa")
	}

	outraFuncao()
	f := funcional{"teste"}
	f.Imprimir("argumento espetacular")
	return nil
}

func (f *funcional) Imprimir(arg string) {
	fmt.Printf("argumento: %s\n", arg)
	fmt.Printf("funciona!: %s\n", f.string)
}
