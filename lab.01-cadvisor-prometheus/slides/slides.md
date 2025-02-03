---
theme: slidev-theme-optiop-theme
colorSchema: dark
author: Mehrshad Lotfi
background: /background.png
layout: cover
title: cAdvisor and Prometheus
info: In this lab, we will monitor the container status using cAdvisor and
  Prometheus. The lab is based on `docker`.
class: text-center
transition: slide-left
favicon: /favicon.png
mdc: true
---

# Metrics with cAdvisor and Prometheus

## Grafana Alloy Crash Course

<br /> 

### Mehrshad Lotfi - DevOps Engineer @ Optiop<span class="color-blue">.</span>

---

## Metrics with cAdvisor and Prometheus

<br />

![cAdvisor Metrics](/cadvisor.prometheus.drawio.svg)

--- 
transition: slide-up
---

## cAdvisor to expose container metrics

![cAdvisor demo](/cadvisor.demo.drawio.svg)



---
transition: slide-up
layout: main
---

# Expose Container Metrics cAdvisor

<br />

````md magic-move {lines: true}
```bash
docker run -d -p 8080:8080 \
  gcr.io/cadvisor/cadvisor
```

```bash
docker run -d -p 8080:8080 --rm \
  --name=lab.01-cadvisor \
  gcr.io/cadvisor/cadvisor
```

```bash {3-7|*}
docker run -d -p 8080:8080 --privileged --rm \
  --name=lab.01-cadvisor \
  --volume=/:/rootfs:ro \
  --volume=/var/run/docker.sock:/var/run/docker.sock:rw \
  --volume=/sys:/sys:ro \
  --volume=/var/lib/docker/:/var/lib/docker:ro \
  --volume=/dev/disk/:/dev/disk:ro \
  gcr.io/cadvisor/cadvisor
```
````
<div class="abs-b m-6">
  <div class="columns-5">
    <LabCheck testURL="http://localhost:1234/api/v1/lab01/test_cadvisor_docker_up" />
  </div>
</div>



---
transition: slide-left
---

## Explore cAdvisor Metrics

<br />


<iframe src="http://localhost:8080/" class="w-full h-4/5" />

[`http://localhost:8080/`](http://localhost:8080/)

<div class="abs-tr m-6">
  <FullscreenButton />
</div>

--- 
transition: slide-up
---

## Prometheus to scrape cAdvisor Metrics

![Prometheus demo](/prometheus.demo.drawio.svg)

---
level: 2
transition: slide-up
---

# Prometheus Configuration for cAdvisor Metrics

<br />

```bash
cat <<EOF > prometheus.yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'cadvisor'
    static_configs:
      - targets: ['lab.01-cadvisor:8080']
EOF
```

---
transition: slide-up
layout: main
---

# Start Prometheus using Docker

<br />

````md magic-move {lines: true}
```bash
docker run -d -p 9090:9090 \
  prom/prometheus:latest
```

```bash
docker run -d -p 9090:9090 --rm --name lab.01-prometheus \
  prom/prometheus:latest
```

```bash
docker run -d -p 9090:9090 --rm --name lab.01-prometheus \
  --volume ./prometheus.yaml:/etc/prometheus/prometheus.yaml \
  prom/prometheus:latest
```

```bash
docker run -d -p 9090:9090 --rm --name lab.01-prometheus \
  --volume ./prometheus.yaml:/etc/prometheus/prometheus.yaml \
  prom/prometheus:latest \
  --config.file=/etc/prometheus/prometheus.yaml
```
````
<div class="abs-b m-6">
<div class="columns-5">
  <LabCheck testURL="http://localhost:1234/api/v1/lab01/test_prometheus_target_exist" />
</div>
</div>

---
transition: slide-left
---

## Explore cAdvisor Metrics

<br />


<iframe src="http://localhost:9090/" class="w-full h-4/5" />

[`http://localhost:9090/`](http://localhost:8080/)

<div class="abs-tr m-6">
  <FullscreenButton />
</div>

--- 
transition: slide-up
---

## Prometheus to scrape cAdvisor Metrics

![Prometheus demo](/grafana.demo.drawio.svg)

--- 
transition: slide-up
---

## Run Grafana using Docker

<br /> 

```bash
docker run -d -p 3000:3000 --rm \
  --name lab.01-grafana grafana/grafana
```

<div class="abs-b m-6">
<div class="columns-5">
  <LabCheck  testURL="http://localhost:1234/api/v1/lab01/test_grafana_reachable" />
</div>
</div>



--- 
transition: slide-left
---

## Explore Grafana

<br />


<iframe src="http://localhost:3000/" class="w-full h-4/5" />

[`http://localhost:3000/`](http://localhost:3000/)

<div class="abs-tr m-6">
  <FullscreenButton />
</div>

--- 
transition: slide-left
---

## Next in Grafana Alloy Crash Course

<br />

![Alloy cAdvisor Prometheus](/alloy.cadvisor.prometheus.drawio.svg)

---
layout: center
class: text-center
---

# Follow Us

<mdi-youtube class="color-red"/> [Youtube](https://www.youtube.com/@Optiop-Group) · 
<mdi-github class="color-white" /> [GitHub](https://github.com/optiop/) · 
<mdi-linkedin class="color-white" /> [Linkedin](https://linkedin.com/company/optiop-group)

<div class="abs-br m-10">
  <h2 class="text-2xl">Optiop <span class="color-blue">.</span></h2>
</div>
