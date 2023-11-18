#!/bin/bash

if [[ -z "$1" ]]; then
    echo "Provide name of the test. Example: ./run.sh microservices_1"
    exit 1;
fi

mkdir results/$1

# Run test
k6 run -vu 10 --iterations 100 --out json="results/$1/data_$(date +"%F_%T").json"  src/test.js
mv results/summary.json "results/$1/summary_$(date +"%F_%T").json"