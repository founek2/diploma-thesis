#!/bin/bash

if [[ -z "$1" ]]; then
    echo "Provide name of the test. Example: ./run.sh microservices_1"
    exit 1;
fi

dir_name="$1"
mkdir -p results-latency/$1

# Run test
k6 run -vu 10 --iterations 100 --out json="results-latency/$1/data_$(date +"%F_%T").json"  src/test_latency.js
mv results-latency/summary.json "results-latency/$1/summary_$(date +"%F_%T").json"