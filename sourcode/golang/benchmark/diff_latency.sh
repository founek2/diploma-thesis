#!/bin/bash


for test_name in "$@"
do
   rps=$(jq '.metrics.http_reqs.values.rate' results-latency/$test_name/summary_*.json)
   med=$(jq '.metrics.http_req_duration.values.med ' results-latency/$test_name/summary_*.json)
   p90=$(jq '.metrics.http_req_duration.values."p(90)" ' results-latency/$test_name/summary_*.json)
   p95=$(jq '.metrics.http_req_duration.values."p(95)" ' results-latency/$test_name/summary_*.json)

   rps=$(printf '%.*f\n' 0 $rps)
   med=$(printf '%.*f\n' 0 $med)
   p90=$(printf '%.*f\n' 0 $p90)
   p95=$(printf '%.*f\n' 0 $p95)

   echo "TEST --- $test_name ----"
   echo "rps=$rps"  
   echo "med=${med}"
   echo "p90=${p90}"
   echo "p95=${p95}"
   echo "latex= " "$rps\,req/s & $med\,ms & $p90\,ms & $p95\,ms" 
   echo ""
done
