# Go Best Practices

The best practices included here comes from the [Golang Worst Practices](https://www.youtube.com/watch?v=Jns0QgJtAYY) 
video by Oliver Powell.

Instead of talking about best practices, it's better to know what are the practices
that we need to avoid in Go.

Here are the things we should avoid:

## 1. Large Interfaces 
The bigger the interface, the weaker the abstraction.

Interfaces have the following characteristics.
1. Readability and understandability. 
2. The best modules are deep. they allow a lot of functionality to be accessed
through a simple interface. A shallow modeule , which is an interface with a lot of 
methods or functions but relatively low functionality.

## 2. Interface Pollution

Abstraction shouldn't be vague. It should create a new semantic
level in which one can be absolutely precise.

Why is this bad?
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

## 3. Excessive mocking
* But don't I need that interface to mock?

*"If you have 100% test coverage and you use mock, no you don't :D"*

Test against the real thing. Test against the real postgreSQL DB if
you need to.

The implementor (not the package) can define the interface they need.

## 4. Over use of DRY
* Anything should have tradeoffs.
* Excessive use of dry might result to entangled code dependencies. If you create a 
generic code instead, the code will bloat, and 
it's now harder to modify, less performance and hard to understand.

How much to dry
* Rule of 3 is a good default.
* More in flux codebase is = more examples needed imo.

## 5. No timouts
* Need timeouts

## 6. Not closing things
* Database rows.
* Files.
* If you open something, make sure it gets closed.

It avoid memory leaks then suddenly your application gets out of memory.

## 7. Bad Naming
* vague, calling something 'data'
* wrong, eg naming a shimpment an order.

## 8. Inconsistency
* naming: database, db, storage, repository.
    - If you call something a db, call it the same thing everywhere.
* style: use the same style/libraries/etc.
    - If it's ugly, at least it's consistently ugly. Rely on a particular
    style when you're writing a code.

## 9. Overreaching concurrencies
* Make sure if we're adding concurrencies, make some benchmarks to see if 
the additional complexities are really worth it.
* Concurrency is hard. CSP doesn't change that.

Don't forget how much work a single thread (or go routine) can do.

Paper: Scalability! But at what COST?