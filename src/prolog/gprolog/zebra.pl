% https://www.101computing.net/solving-a-zebra-puzzle-using-prolog/

/* Facts */
kid(ethan).
kid(ali).
kid(anya).

hero(spiderman).
hero(captain_america).
hero(iron_man).

age(six).
age(eight).
age(ten).

/* Rules */

relation(K,H,A) :- K=anya, H=spiderman, age(A).
relation(K,H,A) :- K=ethan, hero(H), age(A), H\=captain_america.
relation(K,H,A) :- kid(K), H=spiderman, A=six.
relation(K,H,A) :- kid(K), H=captain_america, A=eight.

different(X,Y,Z) :- X\=Y,X\=Z,Y\=Z.
solve(K1,H1,A1,K2,H2,A2,K3,H3,A3) :- relation(K1,H1,A1),relation(K2,H2,A2),relation(K3,H3,A3),different(K1,K2,K3),different(H1,H2,H3),different(A1,A2,A3).
