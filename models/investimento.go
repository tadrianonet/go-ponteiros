package models

type Usuario struct {
	Nome  string
	Idade int
}

type Investimento struct {
	Valor   float64
	Usuario Usuario
}

func AplicarJuros(i *Investimento, taxa float64) {
	i.Valor += i.Valor * taxa / 100
}
