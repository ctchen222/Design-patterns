package decorator

import "fmt"

type Coffeeer interface {
	Cost() int
	Description() string
}

type NormalCoffee struct{}

func (n *NormalCoffee) Cost() int {
	price := 60
	fmt.Printf("Coffee costs %d\n", price)
	return price
}

func (n *NormalCoffee) Description() string {
	return "Coffee"
}

type MilkDecorator struct {
	Coffee Coffeeer
}

func (m *MilkDecorator) Cost() int {
	price := m.Coffee.Cost() + 10
	fmt.Println("Add Milk -> Add 10 NTD")
	fmt.Printf("Latte costs %d\n", price)
	return price
}

func (m *MilkDecorator) Description() string {
	description := m.Coffee.Description() + " + milk"
	return description
}

type SugarDecorator struct {
	Coffee Coffeeer
}

func (m *SugarDecorator) Cost() int {
	price := m.Coffee.Cost() + 5
	fmt.Println("Add Sugar -> Add 10 NTD")
	fmt.Printf("Coffee with sugar costs %d\n", price)
	return price
}

func (m *SugarDecorator) Description() string {
	description := m.Coffee.Description() + " + sugar"
	return description
}
