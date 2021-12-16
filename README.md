# Compiling protobuf to golang

``` shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/meterusage.proto
```

# Compiling protobuf to python

``` shell
cd webserver
python -m grpc_tools.protoc -I../proto --python_out=. --grpc_python_out=. ../proto/meterusage.proto
```

# Building/Running Everything 

To build everything, you need `docker-compose` and `docker` installed on your local machine

``` shell
docker-compose up --build -d
```

# Debugging issues and rebuilding

If you run into problems, you can get rid of the build cache and rebuild everything

``` shell
docker-compose down
docker-compose rm -f
docker-compose up --build -d
```

