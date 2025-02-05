---
theme: slidev-theme-optiop-theme
colorSchema: dark
author: Mehrshad Lotfi
background: /background.png
layout: cover
title: Loki and Promtail
info: In this lab, we will monitor the container logs using Loki and Promtail.
  The lab is based on `docker`.
class: text-center
transition: slide-left
favicon: /favicon.png
mdc: true
---

# Logs with Loki and Promtail

## Grafana Alloy Crash Course

<br /> 

### Mehrshad Lotfi - DevOps Engineer @ Optiop<span class="color-blue">.</span>

---

## Logs with Loki and Promtail

<br />

![Loki logs](/promtail.loki.drawio.svg)

--- 
transition: slide-up
---

## Promtail to collect container logs

![Promtail demo](/promtail.loki.drawio.svg)

---
transition: slide-up
layout: main
---

# Collect Container Logs with Promtail

<br />

````md magic-move {lines: true}
```bash
docker run -d --name=promtail -p 9080:9080 \
  -v /var/lib/docker/containers:/var/lib/docker/containers:ro \
  -v /var/log:/var/log:ro \
  -v /etc:/etc:ro \
  grafana/promtail
```
````

<div class="abs-b m-6">
  <div class="columns-5">
    <LabCheck testURL="" />
  </div>
</div>

---
transition: slide-left
---

## Explore Promtail Logs

<br />

<iframe src="http://localhost:9080/" class="w-full h-4/5" />

[`http://localhost:9080/`](http://localhost:9080/)

<div class="abs-tr m-6">
  <FullscreenButton />
</div>

--- 
transition: slide-up
---

## Loki to store container logs

![Loki demo](/promtail.loki.drawio.svg)

---
level: 2
transition: slide-up
---

# Loki Configuration for Promtail Logs

<br />

```bash
loki configuration for promtail logs
```

---
transition: slide-up
layout: main
---

# Start Loki using Docker

<br />

````md magic-move {lines: true}
```bash
docker run -d --name=loki -p 3100:3100 \
grafana/loki
```
````

<div class="abs-b m-6">
  <div class="columns-5">
    <LabCheck testURL="" />
  </div>
</div>

---
transition: slide-left
---

## Explore Loki Logs

<br />

<iframe src="http://localhost:3100/" class="w-full h-4/5" />

[`http://localhost:3100/`](http://localhost:3100/)

<div class="abs-tr m-6">
  <FullscreenButton />
</div>

--- 
transition: slide-up
---

## Run Grafana using Docker

<br /> 

```bash
docker run -d -p 3000:3000 --rm \
  --name lab.03-grafana grafana/grafana
```

<div class="abs-b m-6">
  <div class="columns-5">
    <LabCheck testURL="" />
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
