# Begin
## install
```
export GOPATH=/Users/<username>/go
export PATH=$GOPATH/bin:$PATH
export GOBIN=$GOPATH/bin
```

```
brew install protobuf
go get -u github.com/golang/protobuf/{proto,protoc-gen-go} go get -u google.golang.org/grpc
go get -u google.golang.org/grpc
```

## week2 demo
```
protoc week2/calculate.proto --go_out=paths=source_relative,plugins=grpc:. 
or
protoc week2/calculate.proto --go_out=plugins=grpc:.
```

## week3-4 demo
```
protoc week3/calculate.proto --go_out=plugins=grpc:.

go run week3/server/main.go
go run week3/client/main.go
```

## week4 demo
```
protoc week4/calculate.proto --go_out=plugins=grpc:.

go run week4/server/main.go
go run week4/client/main.go
```