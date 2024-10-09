package main

import (
	"math/rand"
	"time"
	"github.com/fatih/color"
)
const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message string
	success bool
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	color.Cyan("The Pizzaria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of pizza
	var i = 0

	for {
		currentPizza := makePizza(i)
	}
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making Pizza #%d. It will take %d seconds....\n", pizzaNumber, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message: msg,
			success: successs,
		}

		return &p

	}

	return &PizzaOrder {
		pizzaNumber: pizzaNumber,
	}
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <- ch 
}