# Exercise 11 - Logic Programming

If you do not finish during the lecture period, please finish it as homework.


## Exercise 11.1
Install SWI Prolog from [swi-prolog.org](https://www.swi-prolog.org/).

* Windows users: simply download a [nightly build](https://www.swi-prolog.org/download/daily/bin/) or use the docker container
* Mac users: run `brew install swi-prolog`

Create a new file named `food.pl` and put the following content:
```prolog
likes(sam,Food) :- indian(Food), mild(Food).
likes(sam,Food) :- italian(Food).
likes(sam,chips).

indian(curry).
indian(dahl).
indian(tandoori).

mild(dahl).
mild(tandoori).

italian(pizza).
italian(spaghetti).
```

Run `swipl food.pl` and play with the interactive promt:

* `likes(sam,pizza).`
* `likes(sam,chips).`
* `likes(sam,curry).`
* `likes(sam,fish).`
* `likes(sam,X)` (press n for next result)
* `trace. likes(sam,curry).`
* `halt.`

## Exercise 11.2



## Exercise 11.3
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
