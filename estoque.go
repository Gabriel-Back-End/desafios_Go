package main


import "fmt"


func main() {
    
	estoque := map[string] int {"Espada" : 50, "Escudo" : 80, "Poção" : 20}

	
	
	var nome string


	fmt.Println("Digite o Nome do Produto: ")
	
	fmt.Scan(&nome)
	


	resposta := veridicar_prod(nome, estoque)
	
	
	fmt.Println(resposta)
	

}



func veridicar_prod (nome string, lista map[string]int) string {


	preco, ok := lista[nome]
    
    

	if ok {
        
		return fmt.Sprintf("O produto %v está custando R$ %v,00 Reais", nome,preco)
    
	}else{
        
		return"não existe o Produto no estoque"
    
	}


}