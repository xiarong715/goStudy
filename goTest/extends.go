package main

import "log"

// Engine engine
type Engine interface {
	Start()
	Stop()
}

// Car car
type Car struct {
	Engine
}

func main() {
	car := Car{}
	car.GoToWorkIn()
}

// GoToWorkIn go to work in
func (c *Car) GoToWorkIn() {
	c.Start()
	c.Stop()
}

// Start start
func (c *Car) Start() {
	log.Printf("car start!")
}

// Stop stop
func (c *Car) Stop() {
	log.Printf("car stop!")
}
