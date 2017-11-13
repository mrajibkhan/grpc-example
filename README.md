

put go-grpc-example in any of your $GOPATH/src directory



from go-grpc-example directory execute
```$xslt
$ dep ensure
```

build:
```$xslt
go build go-grpc-example/grpc-example
```
this will create executable **grpc-example** in current directory

install:
 
```$xslt
$ go install go-grpc-example/grpc-example
```

this will create executable **grpc-example** in $GOPATH/bin

proto

from go-grpc-example
```$xslt
protoc -I grpc-example/  --go_out=plugins=grpc:grpc-example/catalog grpc-example/*.proto
```
 