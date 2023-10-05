# load tests

This is an intense type of testing, where we want to check api resilience to the load.

It uses [k6](https://k6.io/docs/get-started/installation/) tool to run the tests, make sure you have it installed on your machine. Or just use it via docker.

_note: api has to running_

```
# run test
cd test/load
k6 run -e HOSTNAME=localhost --vus 100 --duration 30s script.js
```

```
# run test via docker
cd test/load
docker run --rm -i grafana/k6 run -e HOSTNAME=host.docker.internal --vus 100 --duration 30s - <script.js
```
