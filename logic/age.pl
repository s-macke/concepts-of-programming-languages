get_age(BirthYear) :-
	Age is 2020 - BirthYear,
	format('Someone born in ~w is age ~w at the end of ~w.~n', [BirthYear, Age, 2020]).
