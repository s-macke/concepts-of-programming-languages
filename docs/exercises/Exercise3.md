# Exercise 3 - OOP in Go

If you do not finish during the lecture period, please finish it as homework.

## Stack 3.1 (Containers)

Write a generic LIFO container (stack) for all types by using `any` as dynamic type by using the OOP approach. 
The stack should have at least two methods:

- Push(object)
- Pop()

Use the array append function for Push and the slices feature for Pop 

## Exercise 3.2 - Interfaces, Polymorphism and Embedding

The image shows a typical UML design with inheritance, aggregation and polymorph methods.

![oo](../img/03-exercise.png "A typical OO design")

Implement this design as close as possible to the design in Go:

- The `Paint()` method should print the names and values of the fields to the console
- Allocate an array of polymorph objects and call Paint() in a loop
