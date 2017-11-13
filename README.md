

put grpc-example in any of your $GOPATH/src directory

```
go get github.com/mrajibkhan/grpc-example
```

Run server

```
cd $GOPATH/src/github.com/mrajibkhan/grpc-example
go run grpc-catalog/catalog/server/main.go 
```

Run Client (in another terminal)
```
cd $GOPATH/src/github.com/mrajibkhan/grpc-example
go run grpc-catalog/catalog/client/client.go 
```

build:
```
go build grpc-example/grpc-catalog
```
this will create executable **grpc-example** in current directory

install:
 
```
$ go install grpc-example/grpc-catalog
```

this will create executable **grpc-catalog** in $GOPATH/bin

proto

from grpc-example
```
protoc -I grpc-catalog/  --go_out=plugins=grpc:grpc-catalog/catalog grpc-catalog/*.proto
```

run server
```

```
 