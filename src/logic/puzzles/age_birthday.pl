% from https://www.brainzilla.com/logic/logic-grid/basic-1/

:- style_check(-singleton).
solution(Persons) :-
    Persons = [
               person(1, Name1, Age1, BirthMonth1),
               person(2, Name2, Age2, BirthMonth2),
               person(3, Name3, Age3, BirthMonth3)
               ],
    member(person(_, peter,  _,     april), Person), % Peter's birthday is in April
    member(person(_, eric,   7,         _), Person), % Eric is 7
    member(person(_, arnold, _, september), Person), % Arnold's birthday is in September
    member(person(_, peter,  8,         _), Person), % Peter is 8
    member(person(_, _,      9,         _), Person), % age 9 must be present
    member(person(_, _,      _,   january), Person), % january must be present
    member(person(1, arnold, _,         _), Person), % arnold is defined to be number 1
    member(person(2, peter,  _,         _), Person). % peter is defined to be number 2
