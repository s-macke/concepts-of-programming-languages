: ?Prime
  dup
  3 do
    dup i mod
    0= if 0 unloop exit then
  loop
  -1
;

11 ?Prime .

( unloop is the same as r> r> drop drop )