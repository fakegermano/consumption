#!/bin/bash

set -e

# delete all data that is in the bucket if any
influx delete \
    --bucket-id ${DOCKER_INFLUXDB_INIT_BUCKET_ID} \
    --start 2000-01-01T00:00:00Z \
    --stop 2022-01-01T00:00:00Z

# import data from csv measurements into influxdb
influx write \
    --bucket-id ${DOCKER_INFLUXDB_INIT_BUCKET_ID} \
    --header "#constant measurement,usage" \
    --header "#datatype dateTime:RFC3339,double" \
    --file docker-entrypoint-initdb.d/meterusage_rfc.csv
