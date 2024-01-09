# Media content for master's thesis - Analysis and Comparison of Application Architecture: Monolith

Media have the following folder structure:

```
├── README.md                       - brief description of the content of the media
├── golang                          - directory of Golang implementation and benchmarks
│   ├── benchmark                   - directory containing benchmark scenarios and results
│   │   ├── README.md
│   │   ├── _media
│   │   ├── diff.sh                 - helper to extract data from benchmark results
│   │   ├── diff_latency.sh         - helper to extract data from benchmark results
│   │   ├── package.json
│   │   ├── results                 - benchmark results for performance scenario
│   │   ├── results-latency         - benchmark results for latency scenario
│   │   ├── run.sh                  - script to run benchmark for performance scenario
│   │   ├── run_latency.sh          - script to run benchmark for latency scenario
│   │   ├── src                     - contains JavaScript scenario definition for K6 tool
│   ├── docker                      - docker-compose definitions for all variations of applications
│   │   ├── microservices
│   │   ├── microservices-no-db
│   │   ├── modulith
│   │   ├── modulith-no-db
│   │   └── monolith
│   ├── go.work
│   ├── go.work.sum
│   ├── server-microservices        - microservices application implementation
│   ├── server-microservices-no-db  - microservices application implementation (DB replaced by sleep)
│   ├── server-modulith             - modulith application implementation
│   ├── server-modulith-no-db       - modulith application implementation  (DB replaced by sleep)
│   └── server-monolith             - monolith application implementation
├── thesis                          - source form of thesis in format LaTeX
└── thesis.pdf                      - thesis in format pdf
```
