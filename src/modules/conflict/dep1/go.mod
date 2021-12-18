module github.com/dep1

go 1.17

replace github.com/mylib => ../V1/

require (
	github.com/mylib v1.0.0
)