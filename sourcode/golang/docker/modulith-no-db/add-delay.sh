#!/bin/bash

docker compose exec server-modulith tc qdisc add dev eth0 root netem delay 16ms