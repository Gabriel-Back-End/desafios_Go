package main
import (
    "fmt"
    "sort"
    "strings"
)

type aluno struct {
    nome string
    nota float64
}

func main() {
    var turma []aluno
    var resposta string
    
    fmt.Println("Olá Professor, vamos la?")
    
    for{
        
        var novaNota float64
        var novoNome string
        
        fmt.Println("Digite o Nome do aluno: ")
        fmt.Scan(&novoNome)
        
        if novoNome == "fim" {
            break
        }
        
        for{
            fmt.Println("Digite a nota do aluno: (0 a 10)")
            fmt.Scan(&novaNota)
            
            if novaNota >= 0 && novaNota <= 10 {
                break
            }else{
                fmt.Println("Erro !!!\n")
                fmt.Println("A nota tem que ser de 0 a 10\n")
            }
            
        }
        novoAluno := aluno{nome: novoNome, nota: novaNota}   
        
        turma = append(turma, novoAluno)
    }
    
    sort.Slice(turma, func(i,j int) bool{
        return turma[i].nota > turma[j].nota
    })
    
    if len(turma) > 0 {
        
        for{
            fmt.Println("Você deseja ver o Ranking das notas? (s/n)")
            fmt.Scan(&resposta)
            
            resposta = strings.ToLower(resposta)
            
            if resposta == "s" {
                imprimirRanking(turma)
                break
            }else if resposta == "n"{
                break
            }else{
                fmt.Println("Digite apenas s ou n")
            }
        }
        
        media := calcularMedia(turma)
        fmt.Printf("\nA Média Total da Turma foi de %.1f",media)
    }else{
        fmt.Println("Nenhum aluno/nota encontrado")
    }
}

func calcularMedia (lista []aluno) float64{
    var notaTotalTurma float64
    
    for _, aluno := range lista {
        notaTotalTurma += aluno.nota
    }
    
    return notaTotalTurma / float64(len(lista))
}

func imprimirRanking (lista []aluno) {
    fmt.Println("\n--- RANKING OFICIAL ---")
    for i, aluno:= range lista{
        fmt.Printf("%dº lugar: %s - Nota: %.1f\n", i+1, aluno.nome, aluno.nota)
    }
}
