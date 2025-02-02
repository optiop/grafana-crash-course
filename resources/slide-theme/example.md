---1
theme: ./
transition: slide-left
---

# Metrics with cAdvisor and Prometheus

In this lab, we will monitor the container status using cAdvisor and Prometheus. The lab is based on docker.

<div class="pt-12">
  <span @click="$slidev.nav.next" class="px-2 py-1 rounded cursor-pointer" flex="~ justify-center items-center gap-2" hover="bg-white bg-opacity-10">
    Press Space for next page <div class="i-carbon:arrow-right inline-block"/>
  </span>
</div>

---2
layout: main
transition: slide-left

---

# Container Metric Visualization

<div class="flex flex-row p-4 justify-between items-center h-100 ">
  <div v-click=1 class="border p-4 w-1/5 rounded-xl font-bold text-center">cAdvisor</div>
  <div v-click=3 class="border  p-4 w-1/5 rounded-xl font-bold text-center">Prometheus</div>
  <div v-click=5 class="border  p-4 w-1/5 rounded-xl font-bold text-center">Grafana</div>
</div>

<arrow v-click=2 x1="245" y1="335" x2="400" y2="335"   color="var(--optiop-primary)" width="2" arrowSize="1" />

<arrow v-click=4 x1="580" y1="335" x2="735" y2="335"   color="var(--optiop-primary)" width="2" arrowSize="1" />

---
layout: main
---

## First thing first, let us start cAdvisor and explore the metrics.

```bash {none|1|2|3-7|8|all}
docker run -d -p 8081:8080 --privileged --rm \
    --name=lab.01-cadvisor \
    --volume=/:/rootfs:ro \
    --volume=/var/run/docker.sock:/var/run/docker.sock:rw \
    --volume=/sys:/sys:ro \
    --volume=/var/lib/docker/:/var/lib/docker:ro \
    --volume=/dev/disk/:/dev/disk:ro \
    gcr.io/cadvisor/cadvisor
```

This command will start cAdvisor and expose the metrics on port 8081. You can access the metrics at http://localhost:8081/metrics. Try to find out the metric container_cpu_usage_seconds_total in the metrics for the current container lab.01-cadvisor.

---4
layout: test

---

# Test: cAdvisor is reachable via port 8081

---5
layout: test

---

# Test: We can retrieve container information

---6
layout: main

---

# Prometheus to scrape the metrics and store them

Now that we have cAdvisor running, let us start Prometheus to scrape the metrics and store them in its time-series database.

We have to create a configuration file for Prometheus. Create a file prometheus.yaml with the following content:

```bash {none|1|2-3|5-9|all}
cat <<EOF > prometheus.yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'cadvisor'
    static_configs:
      - targets: ['lab.01-cadvisor:8080']
EOF
```

---7
layout: main

---

Start Prometheus with the configuration file prometheus.yaml.

```bash {none|1|2|3|4|all}
docker run -d -p 9091:9090 --rm --name lab.01-prometheus \
  --volume ./prometheus.yaml:/etc/prometheus/prometheus.yaml \
  prom/prometheus:latest \
  --config.file=/etc/prometheus/prometheus.yaml
```

---8
layout: test

---

# Test: Prometheus is reachable via port 9091

---9
layout: main

---

# Grafana

Run the following command to start Grafana

```bash
docker run -d -p 3000:3000  --name lab.01-grafana grafana/grafana
```

---
layout: test
url: localhost:8000
---

# Test: Grafana is reachable via port 3000

---11
layout: main

---

# Cleanup

```bash {none|1|2|3|4|all}
docker stop lab.01-cadvisor
docker stop lab.01-grafana
docker stop lab.01-prometheus
rm prometheus.yaml
```

```bash {none|all}
docker compose down
```

<!--
[Documentation](https://sli.dev) / [GitHub Repo](https://github.com/slidevjs/slidev)
-->