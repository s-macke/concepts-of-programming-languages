97 constant achar
122 constant zchar
26 constant nchars

: r13
  achar -
  13 +
  nchars mod
  achar +
;


achar r13 emit
zchar r13 emit
