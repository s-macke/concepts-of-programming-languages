% https://www.brainzilla.com/logic/self-referential-quiz/basic-1/

:- style_check(-singleton).
:- use_module(library(lists)).

% returns the number of elements in a list for question B
count(L, E, N) :- include(=(E), L, L2), length(L2, N).

% just maps numbers to answers, because count() returns a number
mapNumberToAnswer(N, A) :- X=[a,b,c,d], nth0(N, X, A).

isInAcceptedAnwers(N, Answers) :-
    member(answer(N, a), Answers);member(answer(N, b), Answers);member(answer(N, c), Answers);member(answer(N, d), Answers).

solution(Answers) :-
    Answers = [
               answer(1, A1),
               answer(2, A2),
               answer(3, A3)
               ],
    isInAcceptedAnwers(1, Answers),
    isInAcceptedAnwers(2, Answers),
    isInAcceptedAnwers(3, Answers),
    member(answer(1, A2), Answers), % What is the answer to the second question?
    member(answer(2, X), Answers), count([A1, A2, A3], b, N), mapNumberToAnswer(N, X), % How many correct answers are B?
    member(answer(3, Y), Answers), count([A1, A2, A3], a, M), mapNumberToAnswer(M, Y).  % Is there a question with the correct answer A?



