package main

import (
	"fmt"
	"ponteiros/models"
)

func main() {
	nome := "Thiago"
	idade := 39
	usuario := models.Usuario{Nome: nome, Idade: idade}

	valor := 1000.0
	investimento := models.Investimento{Valor: valor, Usuario: usuario}

	fmt.Printf("Valor inicial do investimento: %.2f\n", investimento.Valor)
	fmt.Printf("Usuário responsável: %s\n", investimento.Usuario.Nome)

	models.AplicarJuros(&investimento, 10)

	fmt.Printf("Valor do investimento após aplicar juros para o usuario %s: é de R$: %.2f\n", investimento.Usuario.Nome, investimento.Valor)
}
