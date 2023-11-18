#!/bin/bash

services="
cart
item
order
payment
invoice1
invoice2
invoice3
invoice4
invoice5
invoice6
invoice7
invoice8
"

for service in $services; do
    echo "Running for $service"
    docker compose exec "server-microservices-$service" tc qdisc add dev eth0 root netem delay 16ms
done