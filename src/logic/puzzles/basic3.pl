% https://www.brainzilla.com/logic/logic-grid/basic-3/

:- style_check(-singleton).

solution(Persons) :-
    Persons = [
               person(1, Name1, Years1, Destinations1),
               person(2, Name2, Years2, Destinations2),
               person(3, Name3, Years3, Destinations3),
               person(4, Name4, Years4, Destinations4)
           ],
           member(person(1, amanda,   Y1, london), Persons),
           member(person(2, jack,     Y1,      _), Persons), dif(Y1, 2015),
           member(person(3, mike,      _,     P1), Persons), dif(P1, rio),
           member(person(4, rachel, 2014,      _), Persons),
           member(person(_, P1,        _,  tokyo), Persons), dif(P1, mike), dif(P1, rachel),
           member(person(_, P2,     2016,      _), Persons), member(P2, [jack, mike]),
           member(person(_, _,      2013,      _), Persons),
           member(person(_, _,         _, sydney), Persons).
