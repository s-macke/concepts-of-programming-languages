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

% Wir definieren folgende Datenstruktur [Mann, Wolf, Ziege, Kohl]
% jede der Variablen kann entweder a oder b annehmen für jeweils eine Seite des Flusses
% start :- [a, a, a, a]
% ziel :-  [b, b, b, b]

% Wenn der Mann bei der Ziege oder bei Wolf und Kohl ist, ist der Zustand sicher.
safe([Farmer,_,Ziege,_]) :- Farmer = Ziege.
safe([Farmer,Wolf,_,Kohl]) :- Farmer = Wolf, Farmer = Kohl.

% Definition der gültige Züge
% Es ist möglich von a nach b und zurück zu wechseln.
cross(a,b).
cross(b,a).

% Der Mann fährt immer mit und kann genau einen Passagier oder nichts mitnehmen).
move([X,X,Ziege,Kohl],wolf,[Y,Y,Ziege,Kohl]) :- cross(X,Y).
move([X,Wolf,X,Kohl],ziege,[Y,Wolf,Y,Kohl]) :- cross(X,Y).
move([X,Wolf,Ziege,X],kohl,[Y,Wolf,Ziege,Y]) :- cross(X,Y).
move([X,Wolf,Ziege,Kohl],nichts,[Y,Wolf,Ziege,Kohl]) :- cross(X,Y).

% Rekursive Lösungsdefinition (ähnlich vollständiger Induktion):
% Der Endzustand ist eine gültige Lösung, d.h. also wenn alle den Fluss
% überquert haben und bei Seite b sind.
solution([b,b,b,b],[]).
% Überquerungen, die zu einem sicheren Zustand führen und einer bekannten
% Lösung führen, sind ebenfalls gültige Lösungen.
solution(State,[Move|OtherMoves]) :- move(State,Move,NextState),
                                     safe(NextState),
                                     solution(NextState,OtherMoves).

run :- length(X,7), solution([a,a,a,a],X), write(X), nl.
% solution([a,a,a,a], [ziege,nichts,wolf,ziege,kohl,nichts,ziege]).
