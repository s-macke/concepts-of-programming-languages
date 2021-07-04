grammar Bool;

// Tokens
NOT: '!';
AND: '&';
OR: '|';
P_OPEN: '(';
P_CLOSE: ')';
VAR: [a-zA-Z0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
expr: NOT expr            # Not
    | expr AND expr       # And
    | expr OR expr        # Or
    | VAR                 # Variable
    | P_OPEN expr P_CLOSE # Parenthesis
    ;

// antlr -Dlanguage=G Bool.g4
