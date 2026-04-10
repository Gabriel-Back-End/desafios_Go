package main

import "fmt"

func main() {
    
    notas := [] int {10,5,7,8,5,6,0,4,6,8}
    qtd_aprovados := 0
    qtd_reprovados := 0
    qtd_notas := len(notas) - 1
    
    for i := 0; i <= qtd_notas; i ++ {
    	if ( notas[i] >= 6){
        	fmt.Printf("Nota %v: Aprovada\n",i)
            qtd_aprovados ++
        }else{
        	fmt.Printf("Nota %v: Reprovada\n",i)
            qtd_reprovados ++
        }
    }
    
    fmt.Printf("Foram Aprovados %v Alunos\n",qtd_aprovados)
    fmt.Printf("Foram Reprovados %v Alunos\n",qtd_reprovados)
    
}