# Demo of grpc pub/sub
```
protoc --go_out=plugins=grpc:.  hello.proto

go run server.go

go run sub.gp

go run pub.go
```