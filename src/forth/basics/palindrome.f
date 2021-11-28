: getfirst   over c@ ;
: getlast    >r 2dup + 1- c@ r> swap ;
: ?palindrome ( c-addr u -- f )
  begin
    dup 1 <=      if 2drop true  exit then
    getfirst getlast <> if 2drop false exit then
    1 /string 1-
  again ;

s" otto" ?palindrome?