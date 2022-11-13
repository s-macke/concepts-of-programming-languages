% https://www.youtube.com/watch?v=gJOZZvYijqk

likes(dan, sally).
likes(sally, dan).
likes(josh, brittney).

dating(X,Y) :- likes(X,Y), likes(Y,X).
%friendship(X,Y) :- likes(X,Y), \+ dating(X,Y).
friendship(X,Y) :- likes(X,Y); likes(Y,X).

print_all([]).
print_all([X|Rest]) :- write(X), nl, print_all(Rest).
% use it for print_all([X, Y]),

q :- likes(X, Y),
     write(X), write(' likes '), write(Y), nl, fail.

:-	initialization(q).
