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

# Intorduction

## Grafana Alloy Crash Course

<br /> 

### Mehrshad Lotfi - DevOps Engineer @ Optiop<span class="color-blue">.</span>

---

# Objectives

<br />

<div class="pl-5">
  <div v-click>
    <h2>
    <mdi-check-circle class="color-green" /> Setting up metrics collection 
    </h2>
  </div>
  <br />
  <div v-click>
    <h2><mdi-check-circle class="color-green" /> Gathering logs</h2>
  </div>
  <br />
  <div v-click>
    <h2>
    <mdi-check-circle class="color-green" /> Setting up alerts</h2>
  </div>
  <br />
  <div v-click>
    <h2>
    <mdi-check-circle class="color-green" />
    Visualising metrics and logs</h2>
  </div>
</div>

---

# Requirements

<br />

<div class="pl-5">
  <div v-click>
    <h2>
      <mdi-checkbox-marked class="color-green" /> Unix command line
    </h2>
  </div>
  <br />
  <div v-click>
    <h2><mdi-checkbox-marked class="color-green" /> Docker
    </h2>
  </div>
</div>

<br />
<hr />
<br />

# Audience
<br />
<div class="pl-5">
  <div v-click>
    <h2>
      <mdi-account-circle class="color-blue" /> System administrators
    </h2>
  </div>
  <br />
  <div v-click>
    <h2><mdi-account-circle class="color-blue" /> Software engineers
    </h2>
  </div>
</div>

---

# Course Structure

<br />

```bash
├── LICENSE
├── Makefile
├── README.md
└── lab.01-cadvisor-prometheus
    ├── README.md
    ├── slides
    ├── exercise
    └── solution
```

---

# Pillars of Monitoring

<br />

<div class="pl-5">
  <div v-click>
    <h2>
    <mdi-magnify class="color-green" /> Profiling
    </h2>
  </div>
  <br />
  <div v-click>
    <h2><mdi-routes class="color-green" /> Tracing</h2>
  </div>
  <br />
  <div v-click>
    <h2>
    <mdi-file-document-multiple-outline class="color-green" /> Logging</h2>
  </div>
  <br />
  <div v-click>
    <h2>
    <mdi-timetable class="color-green" /> Metrics</h2>
  </div>
</div>


--- 
transition: slide-left
---

## Next in Grafana Alloy Crash Course

<br />

![Alloy cAdvisor Prometheus](/cadvisor.prometheus.drawio.svg)

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

<!-- 
Have you ever thought about monitoring your application's performance in real-time, but
you don't know where to begin? Do you have already a monitoring system in place, but
you don't know how to use it effectively?

Hi, I am Mehrshad Lotfi, managing director at Optiop. 
Together with my team, we've setup monitoring systems for many companies, 
where terabytes of data are generated every day. 

We also observe applications we develop for our clients.

For this course, I've taken all those experiences and turned them into a comprehensive
guide to monitor your application using Open Source tools provided by Grafana. 

I cover topic such as, setting up the metrics collection, gathering logs, setting up alerts, 
and visualizing the data. If you have a thorough understanding of these basics of monitoring, 
you can improve your observability and make your application more reliable. 

So let's start monitoring your application together.


Before we get started let’s go over a few assumption about 
the background knowledge needed to get the most out of this course. 

First we would be using Unix command line extensively in this course, 
to install, configure and run applications. 
Therefore, being comfortable with command line is required. 

Also we will setup and familiarise ourselves with Grafana, Alloy, 
Prometheus, Mimir and other open source monitoring stack tool 
and explore the features and capabilities of these using Docker. 
For that reason, having an up to date version of Docker installed 
on your machine is also a requirement. 

I am using macOS during the labs, the same instruction can be applied to UNIX based 
environment as well as WSL environment.


You have access to all exercise files for this course, so that you can follow 
alongside and I highly recommend you do. There is no better way to understand 
a tool but rather hands on and practice using it. 

Exercise files are devided to directories, one for each lab or session. 
Each lab directory has a solution directory which shows the end results, 
exercise directory which has uncompleted exercise files and a slides directory
which contains instruction which we follow throughout the session. 

-->