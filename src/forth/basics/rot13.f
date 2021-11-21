: r13 ( c -- o )
  dup 32 or                                    \ tolower
  dup [char] a [char] z 1+ within if
    [char] m > if -13 else 13 then +
  else drop then ;


: rot13
0 do
    dup i +
    dup
    c@
    r13
    swap
    c!
loop
drop
;


s" This is a string"
2dup
.s
rot13
.s

cr type
