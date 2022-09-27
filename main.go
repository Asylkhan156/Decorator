package main

import "fmt"

type Pizza interface {
	getPrice() int
}

type mexicano struct {
}

func (p *mexicano) getPrice() int {
	return 40
}

type boloneze struct {
}

func (p *boloneze) getPrice() int {
	return 50
}

type tomatoTopping struct {
	Pizza Pizza
}

func (c *tomatoTopping) getPrice() int {
	PizzaPrice := c.Pizza.getPrice()
	return PizzaPrice + 5
}

type cheeseTopping struct {
	Pizza Pizza
}

func (c *cheeseTopping) getPrice() int {
	PizzaPrice := c.Pizza.getPrice()
	return PizzaPrice + 10
}

func main() {
	mexicano := &mexicano{}

	mexicanoWithCheese := &cheeseTopping{Pizza: mexicano}

	mexicanoWithCheeseAndTomato := &tomatoTopping{Pizza: mexicanoWithCheese}

	fmt.Println("Price of Mexicano pizza with cheese and tomato topping is ", mexicanoWithCheeseAndTomato.getPrice())

	boloneze := &boloneze{}
	bolonezeWithTomato := &tomatoTopping{Pizza: boloneze}
	bolonezeWithTomatoAndCheese := &cheeseTopping{Pizza: bolonezeWithTomato}

	fmt.Println("Price of Boloneze pizza with cheese and tomato topping is", bolonezeWithTomatoAndCheese.getPrice())
}
