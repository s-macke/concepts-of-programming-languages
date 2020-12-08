# Exercise 11 - Logic Programming

If you do not finish during the lecture period, please finish it as homework.


## Exercise 11.1 - Warmup
Install SWI Prolog from [swi-prolog.org](https://www.swi-prolog.org/).

* Windows users: simply download a [nightly build](https://www.swi-prolog.org/download/daily/bin/) or use the docker container
* Mac users: run `brew install swi-prolog`

Create a new file named `family.pl` and put the following content:
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

Run `swipl family.pl` and play with the interactive prompt:

* `male(charlie).`
* `female(charlie).`
* `parent(X, charlie).` (press ; or n for next result)
* Write a query to find all grandparents of Charlie
* Write a query to find only grandmothers of Charlie
* Write a query to find all siblings of Bob
* `trace.`, then repeat the last query
* Try to query `X.` to print ALL THINGS POSSIBLE IN THE UNIVERSE!
* Use `halt.` or Ctrl+D (EOF) to exit the interactive prompt


## Exercise 11.2 - Palindrome
Write a predicate `palindrome_list/1` that checks if a list of atoms is a palindrome.  
Hint: have look at the [reverse/2](https://www.swi-prolog.org/pldoc/doc_for?object=reverse/2) predicate.

Write a predicate `palindrome/1` that checks if an atom (or string) is a palindrome.  
Hint: have a look at the [name/2](https://www.swi-prolog.org/pldoc/man?predicate=atom_codes/2), [atom_codes/2](https://www.swi-prolog.org/pldoc/man?predicate=atom_codes/2) and [string_codes/2](https://www.swi-prolog.org/pldoc/man?predicate=string_codes/2) predicates.


## Exercise 11.3 - Solve a Logical
Write a prolog program that derives a solution for the following logical:

> Once upon a time a farmer went to a market and purchased a wolf, a goat, and a cabbage. On his way home, the farmer came to the bank of a river and rented a boat. But crossing the river by boat, the farmer could carry only himself and a single one of his purchases: the wolf, the goat, or the cabbage.
>
> If left unattended together, the wolf would eat the goat, or the goat would eat the cabbage.
>
> The farmer's challenge was to carry himself and his purchases to the far bank of the river, leaving each purchase intact. How did he do it?
>
> [Source](https://en.wikipedia.org/wiki/)

Hint 1: The problem is very similar to the one in [cannibals.pl](../../logic/cannibals.pl).  
Hint 2: Only 7 river crossings are needed for solving the puzzle.
