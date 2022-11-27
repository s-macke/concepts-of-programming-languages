% From: https://www.brainzilla.com/logic/logic-grid/superheroes/

% 3 Kids: Bryan, Sean, Tony
% 3 Ages: 6 years, 8 years, 10 years
% 3 superheroes: Batman, Spiderman, Superman

% Solve this logic puzzle to find out the name, the age and the favorite superhero of each kid.
% 1. Bryan likes Spiderman.
% 2. Tony doesn't like Superman.
% 3. The youngest kid likes Spiderman.
% 4. The kid who likes Superman is 8.

:- style_check(-singleton).

solution(Kids) :-
    Kids = [
               kid(1, Name1, Age1, Superhero1),
               kid(2, Name2, Age2, Superhero2),
               kid(3, Name3, Age3, Superhero3)
               ],
        member(kid(1, bryan, _, spiderman), Kids), % Bryan likes spiderman
        member(kid(3, tony,  _,        X ), Kids), dif(X, superman), % Tony doesn't like superman
        member(kid(_, _,     6, spiderman), Kids), % The youngest kid likes Spiderman.
        member(kid(_, _,     8,  superman), Kids), % The kid who likes Superman is 8.
        member(kid(_, _,    10,         _), Kids),
        member(kid(_, _,     _,    batman), Kids),
        member(kid(_, sean,  _,         _), Kids).