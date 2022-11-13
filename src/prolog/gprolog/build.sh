set -e

# ---
#gprolog --consult-file hello.pl
#gprolog --consult-file hello.pl --query-goal "hello, halt."
# ---

#gprolog --consult-file input.pl
#gplc --no-top-level  input.pl
#./input

# ---

#gprolog --consult-file facts.pl
#gplc --no-top-level  facts.pl

# ---

gplc --no-top-level  reverse.pl
./reverse

# ---


#gplc --no-top-level  hello.pl
#./hello

#gplc --no-top-level nqueens.pl
