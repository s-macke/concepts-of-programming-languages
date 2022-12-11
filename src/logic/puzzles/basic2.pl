% https://www.brainzilla.com/logic/logic-grid/basic-2/

:- style_check(-singleton).

solution(Kids) :-
    Kids = [
               kid(1, Name1, Pets1, Colors1),
               kid(2, Name2, Pets2, Colors2),
               kid(3, Name3, Pets3, Colors3)
           ],
        member(kid(1, angela, _, _), Kids),
        member(kid(2, lisa, fish, X), Kids), dif(X, green),
        member(kid(3, susan, _, red), Kids),
        member(kid(_, _, dog, green), Kids),
        member(kid(_, _, _, blue), Kids),
        member(kid(_, _, _, green), Kids),
        member(kid(_, _, _, red), Kids),
        member(kid(_, _, cat, _), Kids).