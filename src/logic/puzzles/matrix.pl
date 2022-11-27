% https://www.brainzilla.com/logic/logic-equations/

:- style_check(-singleton).
:- use_module(library(lists)).

solution(Answers) :-
    Answers = [
               answer(1, A),
               answer(2, B),
               answer(3, C),
               answer(4, D)
               ],
               member(A, [1,2,3,4]),
    		   member(B, [1,2,3,4]),
    		   member(C, [1,2,3,4]),
               member(D, [1,2,3,4]),
               dif(A, B), dif(A, C), dif(A, D),
               dif(B, C), dif(B, D),
               dif(C, D),
               B + D < 4,
               6 is A * B.
