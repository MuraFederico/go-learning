package main

import "fmt"

type person struct {
	fristName string
	lastName  string
}

type iShuffle interface {
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newName string) {
	p.fristName = newName // do we need to deference p? - no
}

func main() {
	person := person{fristName: "Federico", lastName: "Mura"}
	person.updateName("Bob")
	person.print()
}
