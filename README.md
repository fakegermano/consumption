# Meter Usage project

This is a microservice-based application infrastructure to allow users to access the electricity usage from a set o meters.

It uses the following general architecture and protocols:

``` plantuml
@startuml
interface FLUX as f
interface gRPC as g
interface HTTP as h

database "influxdb" {
[influxdb] as DB
[csv data] as CSV
}
[dataserver] as DATA
[webserver] as WEB
[frontend] as FRONT

user -> FRONT : interface
FRONT -> h : json
h -> WEB : json
WEB -> g : grpc
g -> DATA : grpc
DATA -> f : bolt
f -> DB : bolt
DB -> DATA : line-proto
DATA -> WEB : grpc
WEB -> FRONT : json

CSV ..> DB : init
@enduml
```

## Compiling `protobuf` to golang

To run this, you'll need the `protoc` binaries installed and the appropriate go packages.
``` shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/meterusage.proto
```

## Compiling `protobuf` to python

To run this, you'll need the `grpcio-tools` package installed on python.
``` shell
cd webserver
python -m grpc_tools.protoc -I../proto --python_out=. --grpc_python_out=. ../proto/meterusage.proto
```

## Building/Running Everything 

To build everything, you need `docker-compose` and `docker` installed on your local machine:

``` shell
docker-compose up --build -d
```

## Debugging issues and rebuilding

If you run into problems, you can get rid of the build cache and rebuild everything:

``` shell
docker-compose down
docker-compose rm -f
docker-compose up --build -d
```

## Accessing the running application

After building/running the application is available via the address `localhost:8080` on your browser.
