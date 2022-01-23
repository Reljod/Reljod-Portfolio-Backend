package main

import "fmt"

/*
Large Interfaces - the bigger the interface, the weaker the abstraction.

1. Readability and understandability.
2. The best modules are deep. they allow a lot of functionality to be accessed
through a simple interface. A shallow modeule , which is an interface with a lot of
methods or functions but relatively low functionality.
*/

/*
Interface Pollution.

Abstraction shouldn't be vague. It should create a new semantic
level in which one can be absolutely precise.

why is this bad?
* Not enough examples and good abstraction is hard.
* Forcing your interface onto the implementor.
* Not decoupling anything nor more precise.

Take advantage of Go's implicit interfaces
* Defer hard design decisions.

When do you need an interface?
* When it can have multiple implementations.
* Decoupling from expected change.
* User defined behaviour.

Use interfaces only when you really need it.
Don't pollute your code with interfaces that are not really needed.
*/

/*
Excessive mocking
* But don't I need that interface to mock?

"If you have 100% test coverage and you use mock, no you don't :D"

Test against the real thing. Test againt the real postgreSQL DB if
you need to.

The implementor (not the package) can define the interface they need.

*/

func main() {
	fmt.Println("Best Practices")
}
