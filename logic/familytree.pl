mann(adam).
mann(tobias).
mann(frank).
frau(eva).
frau(daniela).
frau(ulrike).
vater(adam, tobias).
vater(tobias, frank).
vater(tobias, ulrike).
mutter(eva, tobias).
mutter(daniela, frank).
mutter(daniela, ulrike).

grossvater(X, Y) :-
    vater(X, Z),
    vater(Z, Y).

grossvater(X, Y) :-
    vater(X, Z),
    mutter(Z, Y).
