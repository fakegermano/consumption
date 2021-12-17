# Meter Usage project

This is a microservice-based application infrastructure to allow users to access the electricity usage from a set o meters.

It uses the following general architecture and protocols [components.iuml](components.iuml):

![components.iuml](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/fakegermano/consumption/master/components.iuml)

## Component breakdown
This project is structured into 3 main components (and some auxiliary):

- `gRPC` Server / `influxdb` Client: writtern in `golang`, it serves time-series data to clients.
- `gRPC` Client / `API` Server: written in `python`, it calls the previous component via `gRPC` and serves a `http` `API` using `json`
- Frontend(Browser) / `API` Client: written in `React` with `typescript`. It uses `axios` to call the `http` API and display data using `React` and `Bootstrap`

Auxiliary components:
- `influxdb`: A time series database, it acts as the repository server for the `gRPC` server component. All data from the provided `.csv` is loaded into the database on initialisation (using [import.sh](scripts/import.sh))
- `nginx`: Serves as a two-purpose component: to serve built `React` files as static `html/js/css` and as a _reverse-proxy_ to allow the browser to reach the backend components, uses configuration at [nginx.conf](frontend/nginx.conf)
- `convert.py`: one-off script to refactor the input `.csv` in a format `influxdb` understands the time-related data. [convert.py](scripts/convert.py)
- `meterusage.proto`: The protobuf file definition that defines the `gRPC` message formats and services. [meterusage.proto](proto/meterusage.proto)

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
