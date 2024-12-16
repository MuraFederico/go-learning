package main

import "fmt"

type person struct {
	fristName string
	lastName  string
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newName string) {
	(*p).fristName = newName // do we need to *p?
}

func main() {
	person := person{fristName: "Federico", lastName: "Mura"}
	person.updateName("Bob")
	person.print()
}
