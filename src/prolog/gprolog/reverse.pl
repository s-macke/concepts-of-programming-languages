% test if a string is reversed, and do string reversal with it

%reverse([]) --> [].
%reverse([L|Ls]) --> reverse(Ls), [L].
%reverse(Ls, Ls1) :- reverse(Ls, [], Ls1).

myreverse(Xs, Ys) :- myreverse(Xs, [], Ys, Ys).
myreverse([], Ys, Ys, []).
myreverse([X|Xs], Rs, Ys, [_|Bound]) :- myreverse(Xs, [X|Rs], Ys, Bound).

q :- myreverse("Hello", L),
     atom_codes(X, L),
     write(X), nl,
     fail.

:-	initialization(q).
