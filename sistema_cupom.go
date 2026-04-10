package main

import "fmt"

type produto struct {
	nome       string
	preco      float64
	quantidade int
}

func main() {
	carrinho := []produto{
		{nome: "Cadeira", preco: 49.99, quantidade: 4},
		{nome: "Garrafa", preco: 9.99, quantidade: 2},
		{nome: "Celular", preco: 4300, quantidade: 1},
	}

	var cupom string
	var total float64

	total = calculartotal(carrinho)

	fmt.Printf("O total ficou igual a R$ %.2f reais\n\n", total)

	fmt.Println("Insira seu cupom de desconto:")
	fmt.Scan(&cupom)

	if cupom == "DEZ" {
		total *= 0.90 // Reduz o valor em 10%
		fmt.Printf("Cupom Aceito!!!\nAgora o Valor final será de R$ %.2f reais\n", total)
	} else {
		fmt.Printf("Cupom Negado!!!\nO Valor final será de R$ %.2f reais\n", total)
	}

}

func calculartotal(lista []produto) float64 {

	var total float64
	total = 0

	for _, item := range lista {
		// Multiplicação com conversão de tipo necessária no Go
		total += item.preco * float64(item.quantidade)
	}

	return total
}