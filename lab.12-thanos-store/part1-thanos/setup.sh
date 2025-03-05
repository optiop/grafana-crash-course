mkdir -p ./prom-eu1
docker run -i quay.io/thanos/thanosbench:v0.2.0-rc.1 \
  block plan -p continuous-365d-tiny --labels 'cluster="eu1"' --max-time=6h \
  | docker run -v ./prom-eu1:/prom-eu1 -i quay.io/thanos/thanosbench:v0.2.0-rc.1 \
  block gen --output.dir prom-eu1