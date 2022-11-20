% FLUSSÜBERQUERUNG - MANN, WOLF, ZIEGE, KOHL
%
% Ein Mann möchte den Fluss überqueren.
% Dabei kann er nur von a nach b und zurück fahren. Stehen bleiben in einem Zug ist nicht erlaubt.

% Es ist möglich von a nach b und zurück zu wechseln.
cross(a,b).
cross(b,a).

move([X],mann,[Y]) :- cross(X,Y), write('Mann fährt von '), write(X), write(' nach '), write(Y), nl.

solution([b],[]).
solution(State,[Move|OtherMoves]) :- move(State, Move, NextState),solution(NextState,OtherMoves).

run :- length(X,Y), Y < 6,solution([a],X), write(X).


