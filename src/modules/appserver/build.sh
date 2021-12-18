go build -buildmode=plugin service1.go
go build -buildmode=plugin service2.go
go run main.go service.so
