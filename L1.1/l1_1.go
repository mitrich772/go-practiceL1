package main

import (
	"fmt"
)

type Human struct {
	Age  int
	Name string
}

func NewUsusalHuman() Human {
	return Human{Age: 15, Name: "Ivan"}
}
func (h *Human) Meet() {
	fmt.Printf("My name is %s i am %d\n", h.Name, h.Age)
}

type Action struct { //Action со встроенной структурой Human
	Human
}

func (a *Action) MeetAndRun() {
	a.Meet()
	fmt.Println("Bye i need to run!!!")
}

// Переопределение Meet
func (a *Action) Meet() {
	fmt.Printf("Action name is %s i am %d\n", a.Name, a.Age)
}

func main() {
	a := Action{NewUsusalHuman()}
	a.Human.Meet() //Метод Meet(Human)
	a.Meet()       //Метод Meet(Action)
	a.MeetAndRun()
}
