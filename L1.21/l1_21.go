package main

import "fmt"

type PrintableLegacy interface { //неподходящий интерфейс
	printLegacy()
}

type Printable interface { //подходящий интерфейс
	print()
}

type LegacyPrinter struct { //Реализация неподход-го
}

func (lp *LegacyPrinter) printLegacy() {
	fmt.Println("Легаси принт")
}

type Adapter struct { //Реализовываем новый в нем юзаем старую реализацию
	LegacyPrinter
}

func (a *Adapter) print() {
	a.LegacyPrinter.printLegacy() //тут вызываем неподх-юю
}

func useInterface(o Printable) {
	o.print()
}

func main() {
	useInterface(&Adapter{})
}
