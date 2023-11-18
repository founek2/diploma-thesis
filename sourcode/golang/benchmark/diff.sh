#!/bin/bash


for test_name in "$@"
do
   rps=$(jq '.metrics.http_reqs.values.rate' results/$test_name/summary_*.json)
   med=$(jq '.metrics.http_req_duration.values.med ' results/$test_name/summary_*.json)
   p90=$(jq '.metrics.http_req_duration.values."p(90)" ' results/$test_name/summary_*.json)
   p95=$(jq '.metrics.http_req_duration.values."p(95)" ' results/$test_name/summary_*.json)

   echo "TEST --- $test_name ----"
   echo "rps=$rps"  
   echo "med=${med}"
   echo "p90=${p90}"
   echo "p95=${p95}"
   echo ""
done
