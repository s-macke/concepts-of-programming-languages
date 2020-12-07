loves(romeo, juliet).
loves(juliet, romeo) :- loves(romeo, juliet).
loves(william, juliet).

happy(juliet).
with_romeo(juliet).
dances(juliet) :- happy(juliet), with_romeo(juliet).

get_loves_juliet :-
    loves(X, juliet),
    write(X), write(' loves Juliet'), nl.
