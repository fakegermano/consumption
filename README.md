# Compiling protobuf to golang

``` shell
cd proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/meterusage.proto
```

# Compiling protobuf to python

``` shell
cd webserver
python -m grpc_tools.protoc -I../proto --python_out=. --grpc_python_out=. ../proto/meterusage.proto
```

