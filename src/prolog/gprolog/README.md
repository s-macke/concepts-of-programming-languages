# PROLOG

https://www.youtube.com/watch?v=gJOZZvYijqk
https://www.let.rug.nl/bos/lpn//lpnpage.php?pagetype=html&pageid=lpn-htmlse45

| Statement        | Prolog equivalence | 
|------------------|--------------------|
| IF (AND ONLY IF) | :-                 |
| AND              | ,                  |
| OR               | ;                  |
| NOT              | \+                 |
| THEN             | .                  |  ???


max rule:
https://cs.union.edu/~striegnk/learn-prolog-now/html/node89.html

Example Code
https://github.com/didoudiaz/gprolog

even(N):- integer(N), mod(N, 2) =:= 0.

###

sum_list([], 0).
sum_list([H|T], Sum) :-
    sum_list(T, Rest),
    Sum is H + Rest.

###

:- use_module(library(clpfd)).
add(A, B) :- B #= A + 10.
run :- add(X, 20), write(X).

###


###

:- use_module(library(clpfd)).

%A chicken farmer also has some cows for a total of 30 animals, and the animals have 74 legs in all.
%How many chickens does the farmer have?

run :- Chickens + Cows #= 30,
Chickens*2 + Cows*4 #= 74,
Chickens in 0..sup,
Cows in 0..sup,
write(Cows), write(Chickens).
