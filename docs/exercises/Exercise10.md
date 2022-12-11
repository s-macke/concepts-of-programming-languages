# Exercise 10 - Logic Programming

If you do not finish during the lecture period, please finish it as homework.

All exercises can be solved online: [https://swish.swi-prolog.org/](https://swish.swi-prolog.org/) 

Alternatively you can install SWI Prolog from [swi-prolog.org](https://www.swi-prolog.org/).
* Windows users: download the [installer](https://www.swi-prolog.org/download/stable)
  or use the [Docker container](https://hub.docker.com/_/swipl);
  see also [swi-prolog page about Docker](https://www.swi-prolog.org/Docker.html).
* Mac users: run `brew install swi-prolog`

## Exercise 10.1 - Warmup

Create the following rules:
```prolog
male(albert).
male(bob).
male(bill).
male(carl).
male(charlie).
male(dan).
male(edward).

female(alice).
female(betsy).
female(diana).

parent(albert, bob).
parent(albert, betsy).
parent(albert, bill).

parent(alice, bob).
parent(alice, betsy).
parent(alice, bill).

parent(bob, carl).
parent(bob, charlie).
parent(diana, charlie).
```

Play with the interactive prompt:

* `male(charlie).`
* `female(charlie).`
* `parent(X, charlie).`
* Write a query to find all grandparents of Charlie
* Write a query to find only grandmothers of Charlie
* Write a query to find all siblings of Bob
* Add `trace, ....`, then repeat the last query

## Exercise 10.2 - Logical grid puzzles

Solve the following puzzle using Prolog:
https://www.brainzilla.com/logic/logic-grid/basic-2/

Hint 1: The problem is very similar to the one in [superheroes.pl](../../src/logic/puzzles/superheroes.pl).
Hint 2: Use dif/2 to test for inequality. For example, `dif(X, Y)` is true if X and Y are not equal.

All kind of puzzles listed here can be solved with the same patterns: https://www.brainzilla.com/logic/

## Exercise 10.3 - Solve a Logical

Write a prolog program that derives a solution for the following logical:

> Once upon a time a farmer went to a market and purchased a wolf, a goat, and a cabbage. On his way home, the farmer came to the bank of a river and rented a boat. But crossing the river by boat, the farmer could carry only himself and a single one of his purchases: the wolf, the goat, or the cabbage.
>
> If left unattended together, the wolf would eat the goat, or the goat would eat the cabbage.
>
> The farmer's challenge was to carry himself and his purchases to the far bank of the river, leaving each purchase intact. How did he do it?
>
> [Source](https://en.wikipedia.org/wiki/Wolf,_goat_and_cabbage_problem)

Hint 1: The problem is very similar to the one in [cannibals.pl](../../src/logic/cannibals.pl).  
Hint 2: Only 7 river crossings are needed for solving the puzzle.
