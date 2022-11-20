factorial(0, 1).
factorial(N, F) :-
    N > 0,
    Prev is N -1,
    factorial(Prev, R),
    F is R * N.

run :- factorial(5, F), write(F), nl.

%% ----------------------------------------------------------------------

:- use_module(library(clpfd)).
factorial2(0, 1).
factorial2(N, F) :-
        N #> 0,
        N1 #= N - 1,
        F #= N * F1,
        factorial2(N1, F1).

run2 :- factorial2(N, 6), write(N), nl.
