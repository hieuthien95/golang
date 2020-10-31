# Begin

## docs
```
https://grpc.io/docs/languages/go/basics/
```

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

## week3: setup server/client
```
protoc week3/calculate.proto --go_out=plugins=grpc:.

go run week3/server/main.go
go run week3/client/main.go
```

## week4: UNARY
```
protoc week4/calculate.proto --go_out=plugins=grpc:.

go run week4/server/main.go
go run week4/client/main.go
```

## week5: SERVER STREAMING
```
protoc week5/server-streaming.proto --go_out=plugins=grpc:.

go run week5/server/main.go
go run week5/client/main.go
```

## week6: CLIENT STREAMING
```
protoc week6/client-streaming.proto --go_out=plugins=grpc:.

go run week6/server/main.go
go run week6/client/main.go
```

## week6: 2 WAY STREAMING
```
protoc week7/min-max-2-way-streaming.proto.proto --go_out=plugins=grpc:.

go run week7/server/main.go
go run week7/client/main.go
```

