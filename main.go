package main

import (
	"fmt"
)

// Item = item do pedido, com um nome e preco.
type Item struct {
	Nome  string
	Preco float64
}

// Pedido = um pedido, contendo um userID e uma lista de itens.
type Pedido struct {
	UserID int
	Itens  []Item
}

// valorTotal = calcular o valor total dos itens no pedido.
func (p *Pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.Itens {
		total += item.Preco
	}
	return total
}
func main() {
	// Criar alguns itens
	item1 := Item{Nome: "Biscoito", Preco: 5.50}
	item2 := Item{Nome: "Bolacha", Preco: 5.75}

	// Criar um pedido com os itens
	pedido := Pedido{
		UserID: 19,
		Itens:  []Item{item1, item2},
	}
	// Calcular valor total do pedido
	total := pedido.valorTotal()

	fmt.Printf("O valor total do pedido é: %.2f\n", total)
}
