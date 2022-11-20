% FLUSSÜBERQUERUNG - MISSIONARE UND KANNIBALE
%
% Drei Missionare und drei Kannibalen wollen einen Fluss überqueren,
% das Boot bietet aber nur Platz für zwei Personen. Um nicht fürchten
% zu müssen aufgefressen zu werden, dürfen die Missionare niemals
% in Unterzahl gegenüber den Kannibalen sein.

% The river can be crossed from a to b and vice versa
cross(a,b).
cross(b,a).

% The state vector saves the position (a or b) for each actor.
% First entry is the board, followed by missionaries and cannibals.

% Single crossing
move([X,X,M2,M3,K1,K2,K3], m1, [Y,Y,M2,M3,K1,K2,K3]) :- cross(X,Y). % not required for solution
move([X,M1,X,M3,K1,K2,K3], m2, [Y,M1,Y,M3,K1,K2,K3]) :- cross(X,Y). % not required for solution
move([X,M1,M2,X,K1,K2,K3], m3, [Y,M1,M2,Y,K1,K2,K3]) :- cross(X,Y). % not required for solution
move([X,M1,M2,M3,X,K2,K3], k1, [Y,M1,M2,M3,Y,K2,K3]) :- cross(X,Y).
move([X,M1,M2,M3,K1,X,K3], k2, [Y,M1,M2,M3,K1,Y,K3]) :- cross(X,Y).
move([X,M1,M2,M3,K1,K2,X], k3, [Y,M1,M2,M3,K1,K2,Y]) :- cross(X,Y).
% Partner crossing
move([X,X,X,M3,K1,K2,K3], m1m2, [Y,Y,Y,M3,K1,K2,K3]) :- cross(X,Y).
move([X,M1,X,X,K1,K2,K3], m2m3, [Y,M1,Y,Y,K1,K2,K3]) :- cross(X,Y).
move([X,X,M2,X,K1,K2,K3], m1m3, [Y,Y,M2,Y,K1,K2,K3]) :- cross(X,Y).
move([X,M1,M2,M3,X,X,K3], k1k2, [Y,M1,M2,M3,Y,Y,K3]) :- cross(X,Y).
move([X,M1,M2,M3,K1,X,X], k2k3, [Y,M1,M2,M3,K1,Y,Y]) :- cross(X,Y).
move([X,M1,M2,M3,X,K2,X], k1k3, [Y,M1,M2,M3,Y,K2,Y]) :- cross(X,Y).
move([X,X,M2,M3,X,K2,K3], m1k1, [Y,Y,M2,M3,Y,K2,K3]) :- cross(X,Y).
move([X,X,M2,M3,K1,X,K3], m1k2, [Y,Y,M2,M3,K1,Y,K3]) :- cross(X,Y).
move([X,X,M2,M3,K1,K2,X], m1k3, [Y,Y,M2,M3,K1,K2,Y]) :- cross(X,Y).
move([X,M1,X,M3,X,K2,K3], m2k1, [Y,M1,Y,M3,Y,K2,K3]) :- cross(X,Y).
move([X,M1,X,M3,K1,X,K3], m2k2, [Y,M1,Y,M3,K1,Y,K3]) :- cross(X,Y).
move([X,M1,X,M3,K1,K2,X], m2k3, [Y,M1,Y,M3,K1,K2,Y]) :- cross(X,Y).
move([X,M1,M2,X,X,K2,K3], m3k1, [Y,M1,M2,Y,Y,K2,K3]) :- cross(X,Y).
move([X,M1,M2,X,K1,X,K3], m3k2, [Y,M1,M2,Y,K1,Y,K3]) :- cross(X,Y).
move([X,M1,M2,X,K1,K2,X], m3k3, [Y,M1,M2,Y,K1,K2,Y]) :- cross(X,Y).


% No missionaries or missionaries are not short-handed
saferatio(0,_).
saferatio(M,K) :- M>=K.

% Helper function for counting missionaries and cannibals
count(X, [X,X,X], 3).
count(X, [X,X,Y], 2) :- X\=Y.
count(X, [Y,X,X], 2) :- X\=Y.
count(X, [X,Y,X], 2) :- X\=Y.
count(X, [X,Y,Y], 1) :- X\=Y.
count(X, [Y,X,Y], 1) :- X\=Y.
count(X, [Y,Y,X], 1) :- X\=Y.
count(X, [Y,Y,Y], 0) :- X\=Y.

%  A state is safe if the ratio of missionaries and cannibals is safe on both river sides
safe([_,M1,M2,M3,K1,K2,K3]) :- count(a, [M1,M2,M3], MA), count(a, [K1,K2,K3], KA), saferatio(MA,KA),
                               count(b, [M1,M2,M3], MB), count(b, [K1,K2,K3], KB), saferatio(MB,KB).

% Recursive solution, similar to mathematical induction.
% Final state is a valid solution, i.e. when all people have crossed the river to side b.
solution([b,b,b,b,b,b,b], []).
% Crossings that lead to a save state and a known solution are also valid solutions.
solution(State, [Move|OtherMoves]) :- move(State,Move,NextState),
                                      safe(NextState),
                                      solution(NextState,OtherMoves).

% Speed up the recursive calculation by remembering the previous results in a table
:- table solution/2.
% Only search for solutions with 11 river crossings and print it
run :- length(X,11), solution([a,a,a,a,a,a,a], X), write(X), nl.

% Solution in 11 moves:
% a,b = river sides, m = missionary, k = cannibal
%     1     2     3     4     5     6     7     8     9    10    11
%  a---->b---->a---->b---->a---->b---->a---->b---->a---->b---->a---->b
%    kk    k_    kk    k_    mm    mk    mm    k_    kk    k_    kk
% 3,3   3,1   3,2   3,0   3,1   1,1   2,2   0,2   0,3   0,1   0,2   0,0  ~  m,k at a
% 0,0   0,2   0,1   0,3   0,2   2,2   1,1   3,1   3,0   3,2   3,1   3,3  ~  m,k at b

% [B,  M1,M2,M3, K1,K2,K3]
% [a,  a, a, a,  a, a, a ]  k1k2
% [b,  a, a, a,  b, b, a ]  k2
% [a,  a, a, a,  b, a, a ]  k2k3
% [b,  a, a, a,  b, b, b ]  k3
% [a,  a, a, a,  b, b, a ]  m1m2
% [b,  b, b, a,  b, b, a ]  m2k2
% [a,  b, a, a,  b, a, a ]  m2m3
% [b,  b, b, b,  b, a, a ]  k1
% [a,  b, b, b,  a, a, a ]  k1k2
% [b,  b, b, b,  b, b, a ]  k2
% [a,  b, b, b,  b, a, a ]  k2k3
% [b,  b, b, b,  b, b, b ]

% Verify solution:
% solution([a,a,a,a,a,a,a], [k1k2,k2,k2k3,k3,m1m2,m2k2,m2m3,k1,k1k2,k2,k2k3]).

% For a quicker result without tabling:
% solution([a,a,a,a,a,a,a], [k1k2,k2,k2k3,k3,m1m2,m2k2,m2m3,k1,X,Y,Z]).
