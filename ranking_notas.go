
package main
import "fmt"
import "sort"

type aluno struct {
    nome string
    nota float64
}

func main() {
    var turma []aluno
    var notaTotalTurma float64
    var mediaFinal float64
    
    fmt.Println("Olá Professor, vamos la?")
    
    for{
        
        var novaNota float64
        var novoNome string
        
        fmt.Println("Digite o Nome do aluno: ")
        fmt.Scan(&novoNome)
        
        if novoNome == "fim" {
            break
        }
        
        fmt.Println("Digite a nota do aluno: ")
        fmt.Scan(&novaNota)
        
        novoAluno := aluno{nome: novoNome, nota: novaNota}   
        
        turma = append(turma, novoAluno)
    }
    
    sort.Slice(turma, func(i,j int) bool{
        return turma[i].nota > turma[j].nota
    })
    
    for i, aluno:= range turma{
        fmt.Printf("%dº lugar: %s - Nota: %.1f\n", i+1, aluno.nome, aluno.nota)
        notaTotalTurma += aluno.nota
    }
    
    mediaFinal = notaTotalTurma / float64(len(turma))
    
    fmt.Printf("E a Média Total da turma foi de %.1f",mediaFinal)
    
}
