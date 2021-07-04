% Palindrome predicate for lists
palindrome_list(L) :- reverse(L, L).

% Palindrome predicate for atoms and strings
palindrome(X) :-
    name(X, L), % actually the same as atom_codes, but being traditional here
    reverse(L, L).

% Palindrome predicate for atoms only
palindrome_atom(A) :-
    atom(A), % restricts this predicate to atoms (single quoted or no quotes)
    atom_codes(A, L),
    reverse(L, L).

% Palindrome predicate for strings only
palindrome_str(S) :-
    string(S), % restricts this predicate to strings (double quoted)
    string_codes(S, L),
    reverse(L, L).
