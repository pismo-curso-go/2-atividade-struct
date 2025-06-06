package main

import (
	"fmt"
)


type Item struct {
	id      int
	produto string
	valor   float64
}


type Pedido struct {
	userID int
	Itens  []Item
}


func (p Pedido) valorTotal() float64 {
	var total float64
	for _, item := range p.Itens {
		total += item.valor
	}
	return total
}

func main() {

	item1 := Item{id: 1, produto: "Notebook", valor: 3000.00}
	item2 := Item{id: 2, produto: "Mouse", valor: 150.00}
	item3 := Item{id: 3, produto: "Teclado", valor: 200.00}


	pedido := Pedido{
		userID: 101,
		Itens:  []Item{item1, item2, item3},
	}


	fmt.Printf("Valor total do pedido: R$ %.2f\n", pedido.valorTotal())
}
