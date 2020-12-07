# Exercise 11 - Logic Programming

If you do not finish during the lecture period, please finish it as homework.


## Exercise 11.1
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

Run `swipl family.pl` and play with the interactive promt:

* `male(charlie).`
* `female(charlie).`
* `parent(X, charlie).` (press ; or n for next result)
* Write a query to find all grandparents of Charlie
* Write a query to find only grandmothers of Charlie
* Write a query to find all siblings of Bob
* `trace.`, then repeat the last query
* `halt.`


## Exercise 11.2
Write a prolog program that calculates a solution for the following problem.

```prolog
% FLUSSÜBERQUERUNG - MANN, WOLF, ZIEGE, KOHL
%
% Ein Mann möchte zusammen mit einem Wolf, einer Ziege und einem Kohlkopf
% einen Fluss überqueren, doch das Boot kann außer ihm nur einen weiteren
% Passagier fassen. Er kann weder den Wolf mit der Ziege noch die Ziege
% mit dem Kohl unbeaufsichtigt an einem Ufer lassen. Aufgabe ist es,
% einen Plan zu entwickeln, der diese Bedingungen einhält und mit möglichst
% wenigen Überfahrten auskommt.
%
% Hinweis: es sind nur 7 Züge erforderlich.

% TODO YOUR CODE HERE
```

Hint: The problem is very similar to the one in [cannibals.pl](../../logic/cannibals.pl).
