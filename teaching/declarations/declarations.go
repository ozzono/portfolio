package declaracoes

import (
	"fmt"
)

type TipoMagico struct {
	CampoMagico    string // ""
	MedidorDeMagia int    // 0
}

func Declarar() {
	var variavelMagica TipoMagico

	outraMagia := TipoMagico{}

	maisUmaMagia := TipoMagico{
		CampoMagico:    "haja magia",
		MedidorDeMagia: 8000,
	}

	retornoMagico, erroMagico := funcaoMagica()

	// declaração anônima não funciona sem declarar valores, senão seria a declaração de um tipo
	// o tipo não pode ser anônimo, no máximo pode ser interno ao escopo
	// antepenultimaMagia:=struct type{
	// campoMagico string
	// medidorDeMagia int
	// }

	type naoMagico struct {
		Nome           string
		MedidorDeMagia int
	}

	// tipo anônimo
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
		"implicita",
		0,
	}

	var pareiComAMagia = TipoMagico{
		"deixa pra lá",
		0,
	}

	fmt.Println("variavelMagica: ", variavelMagica)
	fmt.Printf("variavelMagica: %#v\n", variavelMagica)
	fmt.Println("outraMagia: ", outraMagia)
	fmt.Printf("outraMagia: %#v\n", outraMagia)
	fmt.Println("maisUmaMagia: ", maisUmaMagia)
	fmt.Printf("maisUmaMagia: %#v\n", maisUmaMagia)
	fmt.Println("retornoMagico: ", retornoMagico)
	fmt.Printf("retornoMagico: %#v\n", retornoMagico)
	fmt.Println("erroMagico: ", erroMagico)
	fmt.Println("penultimaMagia: ", penultimaMagia)
	fmt.Printf("penultimaMagia: %#v\n", penultimaMagia)
	fmt.Println("ultimaMagia: ", ultimaMagia)
	fmt.Printf("ultimaMagia: %#v\n", ultimaMagia)
	fmt.Println("pareiComAMagia: ", pareiComAMagia)
	fmt.Printf("pareiComAMagia: %#v\n", pareiComAMagia)
}

func funcaoMagica() (TipoMagico, error) {
	return TipoMagico{}, nil
}
