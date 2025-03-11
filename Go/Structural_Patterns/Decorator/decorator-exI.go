package decorator

import "fmt"

type Pizzaer interface {
	GetPrice() int
}

type NormalPizza struct{}

func (n *NormalPizza) GetPrice() int {
	price := 169
	fmt.Printf("Normal pizza costs %d\n", price)
	return price
}

type TomatoToppingPizza struct {
	Pizza Pizzaer
}

func (t *TomatoToppingPizza) GetPrice() int {
	price := t.Pizza.GetPrice() + 20
	fmt.Println("Tomato topping -> Add 20 NTD")
	fmt.Printf("Tomato pizza costs %d\n", price)
	return price
}

type CheeseToppingPizza struct {
	Pizza Pizzaer
}

func (c *CheeseToppingPizza) GetPrice() int {
	price := c.Pizza.GetPrice() + 30
	fmt.Println("Cheese topping -> Add 30 NTD")
	fmt.Printf("Cheese pizza costs %d\n", price)
	return price
}
