package main
import "fmt"

type Produto struct{
    ID int
    Nome string
    Preco float64
}

func main() {
    
    estoque := make(map[int]Produto)
    
    var totalCompra float64
    var formaPG int
    
    totalCompra = 0  
    
    itens := []Produto{
        {ID: 1,  Nome: "Arroz 5kg",      Preco: 25.50},
        {ID: 2,  Nome: "Feijão 1kg",     Preco: 8.90},
        {ID: 3,  Nome: "Leite 1L",       Preco: 5.20},
        {ID: 4,  Nome: "Café 500g",      Preco: 14.30},
        {ID: 5,  Nome: "Açúcar 1kg",     Preco: 4.50},
        {ID: 6,  Nome: "Óleo de Soja",   Preco: 7.80},
        {ID: 7,  Nome: "Macarrão 500g",  Preco: 3.90},
        {ID: 8,  Nome: "Bolacha Salgada", Preco: 4.20},
        {ID: 9,  Nome: "Detergente",     Preco: 2.50},
        {ID: 10, Nome: "Sabão em Pó",    Preco: 18.90},
        {ID: 11, Nome: "Papel Higiênico", Preco: 15.00},
        {ID: 12, Nome: "Creme Dental",   Preco: 3.50},
        {ID: 13, Nome: "Shampoo",        Preco: 12.00},
        {ID: 14, Nome: "Sabonete",       Preco: 2.20},
        {ID: 15, Nome: "Manteiga 500g",  Preco: 11.50},
        {ID: 16, Nome: "Pão de Forma",   Preco: 7.00},
        {ID: 17, Nome: "Suco de Uva 1L", Preco: 13.00},
        {ID: 18, Nome: "Ovos (Dúzia)",   Preco: 9.00},
        {ID: 19, Nome: "Frango (kg)",    Preco: 16.50},
        {ID: 20, Nome: "Sal 1kg",        Preco: 2.10},
    }
    
    for _, p := range itens {
        estoque[p.ID] = p
    }
    
    for {
      
        var idDigitado int
        var qtdProd int
        
        fmt.Println("\n\n\n\nDigite o código do produto: ")
        fmt.Scan(&idDigitado)
        
        if idDigitado == 0 {
            if totalCompra > 0{
                break
            }else {
                fmt.Println("Nenhum produto selecionado")        
            }
        }else {
        
            produto, existe := estoque[idDigitado]
            
            if existe {
                fmt.Printf("Quantos %v você deseja ?\n", produto.Nome)
                fmt.Scan(&qtdProd)
                
                totalCompra += (produto.Preco * float64(qtdProd))
                
            }else {
                fmt.Println("Produto Invalido")
            }
        }
    
    }
    
    fmt.Printf("\n\n\nO total da Compra foi de R$ %.2f reais\n\n", totalCompra)
    
    for {
        fmt.Println("Qual sera a forma de pagamento?\n 1 - Débito\n 2 - Crédito\n 3 - Dinheiro")
        fmt.Scan(&formaPG)
        
        if formaPG != 1 && formaPG != 2 && formaPG != 3 {
            fmt.Println("Forma de Pagamento Invalida")
        }else{
            fmt.Println("Forma de Pagamento selecionada ", formaPG)
            break
        }
    }
    
    if formaPG == 3 {
        
        var recebido float64
        var troco float64
        
        for {
            fmt.Println("\n\nQuanto o cliente pagou ?")
            fmt.Scan(&recebido)
            
            if recebido >= totalCompra {
                break
            }else{
                troco = totalCompra - recebido
                
                fmt.Printf("esta faltando R$ %.2f\n", troco)
            }
        }
        
        
        troco = recebido - totalCompra
        
        fmt.Printf("Troco será R$ %.2f \n", troco)
    }
    
    fmt.Println("Obrigado pela preferencia")
    
}
