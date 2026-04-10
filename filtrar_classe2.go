package main

import "fmt"

type item struct {
	Nome   string
	Preco  int
	Classe string
}

func main() {

	estoque := []item{
		{Nome: "Espada", Preco: 50, Classe: "Guerreiro"},
		{Nome: "Arco", Preco: 35, Classe: "Arqueiro"},
		{Nome: "Cajado", Preco: 60, Classe: "Mago"},
		{Nome: "Machado de Batalha", Preco: 55, Classe: "Guerreiro"},
		{Nome: "Besta Pesada", Preco: 45, Classe: "Arqueiro"},
		{Nome: "Grimório Antigo", Preco: 90, Classe: "Mago"},
		{Nome: "Escudo de Carvalho", Preco: 30, Classe: "Guerreiro"},
		{Nome: "Varinha de Fênix", Preco: 75, Classe: "Mago"},
		{Nome: "Aljava de Prata", Preco: 20, Classe: "Arqueiro"},
		{Nome: "Martelo de Guerra", Preco: 65, Classe: "Guerreiro"},
		{Nome: "Esfera de Cristal", Preco: 85, Classe: "Mago"},
		{Nome: "Arco Longo de Elite", Preco: 70, Classe: "Arqueiro"},
		{Nome: "Lança de Ferro", Preco: 40, Classe: "Guerreiro"},
		{Nome: "Capa de Mana", Preco: 50, Classe: "Mago"},
		{Nome: "Flechas de Fogo", Preco: 25, Classe: "Arqueiro"},
	}

	var escolhido int
	var classe string

	fmt.Println("Qual Clase Você está Procurando?")
	fmt.Println("1. Mago")
	fmt.Println("2. Guerreiro")
	fmt.Println("3. Arqueiro")
	fmt.Scan(&escolhido)

	if escolhido == 1 {
		classe = "Mago"
		filtrarporclasse(estoque, classe)
	} else if escolhido == 2 {
		classe = "Guerreiro"
		filtrarporclasse(estoque, classe)
	} else if escolhido == 3 {
		classe = "Arqueiro"
		filtrarporclasse(estoque, classe)
	} else {
		fmt.Println("Classe Não Encontrada")
	}

}

func filtrarporclasse(lista []item, classedesejada string) {
	fmt.Printf("-------------- Itens da Clase %v --------------\n", classedesejada)

	contador := 0
	soma := 0

	for _, item := range lista {

		if item.Classe == classedesejada {
			fmt.Printf("Produto: %v - Preço: R$ %v,00\n", item.Nome, item.Preco)
			contador++
			soma += item.Preco
		}
	}

	if contador > 0 {
		media := float64(soma) / float64(contador)
		fmt.Printf("\n%v Produtos Encontrados\nA média dos Valores é igual a R$ %.2f reais\n", contador, media)
	} else {
		fmt.Println("Nenhum item encontrado para esta classe.")
	}
}