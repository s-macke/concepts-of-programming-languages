: exchange
  2dup c@ swap c@ rot c! swap c! ;

: reverse
2dup + 1- rot rot
2 / 0 do
2dup exchange
1+ swap 1- swap
loop
2drop
;

s" This is a string" 2dup reverse cr type