: printEachChar
0 do
    dup i +
    c@ emit cr
loop
;

s" This is a string" printEachChar