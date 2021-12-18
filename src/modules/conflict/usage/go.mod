module usage

go 1.17

replace github.com/dep1 => ../dep1/
replace github.com/dep2 => ../dep2/

replace github.com/mylib v1.0.0 => ../V1/
replace github.com/mylib v1.0.1 => ../V2/

require github.com/dep1 v1.0.0
require github.com/dep2 v1.0.0
