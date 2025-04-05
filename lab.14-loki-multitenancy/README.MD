This guide walks you through the steps to set up **Grafana Loki with multiple Alloy agents**, each forwarding logs from different sources (containers and system logs) to Loki under **separate tenants**.

---

## 🧩 Architecture

- **Loki**: Central log aggregator.
- **Alloy (1)**: Forwards **container logs** to Loki (tenant: `optiop1`).
- **Alloy (2)**: Forwards **system logs** to Loki (tenant: `optiop2`).
- **Grafana**: Visualizes the logs from Loki using provisioned data sources.

---

## 📁 Folder Structure

```
project-root/
│
├── docker-compose.yml
├── config.alloy        # For Alloy (container logs)
├── config2.alloy       # For Alloy (system logs)
└── local-config.yaml   # For Loki
└── ds.yaml             # For Grafana (data sources)
```

---

## 🐳 Step 1: Create `docker-compose.yml`

```yaml

 services:
   loki:
     image: grafana/loki:3.0.0
     hostname: loki
     volumes:
       - ./local-config.yaml:/etc/loki/local-config.yaml
     networks:
      - app_net
     ports:
       - "3100:3100"
     command: -config.file=/etc/loki/local-config.yaml


   alloy:
    # Grafana Alloy version 1.5.1
    # https://hub.docker.com/r/grafana/alloy/tags
    image: grafana/alloy:v1.7.5
    privileged: true
    ports:
      - 12345:12345
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:rw
      - ./config.alloy:/etc/alloy/config.alloy
    networks:
     - app_net
    command:
      - run
      - /etc/alloy/config.alloy
      - --storage.path=/var/lib/alloy/data
      - --server.http.listen-addr=0.0.0.0:12345
   alloy-2:
    # Grafana Alloy version 1.5.1
    # https://hub.docker.com/r/grafana/alloy/tags
    image: grafana/alloy:v1.7.5
    privileged: true
    networks:
     - app_net
    ports:
      - 12346:12345
    volumes:
      - /var/log:/var/log:ro
      - ./config2.alloy:/etc/alloy/config.alloy
    command:
      - run
      - /etc/alloy/config.alloy
      - --storage.path=/var/lib/alloy/data
      - --server.http.listen-addr=0.0.0.0:12345

   grafana:
     environment:
       - GF_PATHS_PROVISIONING=/etc/grafana/provisioning
       - GF_AUTH_ANONYMOUS_ENABLED=true
       - GF_AUTH_ANONYMOUS_ORG_ROLE=Admin
     networks:
      - app_net
     volumes:
       - ./ds.yaml:/etc/grafana/provisioning/datasources/ds.yaml
     entrypoint:
       - sh
       - -euc
       - /run.sh
     image: grafana/grafana:11.0.0
     ports:
       - "3000:3000"
 networks:
  app_net:
   name: app_net
    
```

---

## 📜 Step 2: Create `config.alloy` for Alloy (Container Logs)

```hcl
discovery.docker "flog_scrape" {
  host = "unix:///var/run/docker.sock"
  refresh_interval = "5s"
}

discovery.relabel "flog_scrape" {
  targets = []

  rule {
    source_labels = ["__meta_docker_container_name"]
    regex         = "/(.*)"
    target_label  = "container"
  }

  rule {
    source_labels = ["__meta_docker_container_log_stream"]
    target_label  = "logstream"
  }

  rule {
    source_labels = ["__meta_docker_container_label_logging_jobname"]
    target_label  = "job"
  }
}

loki.source.docker "flog_scrape" {
  host             = "unix:///var/run/docker.sock"
  targets          = discovery.docker.flog_scrape.targets
  forward_to       = [loki.write.default.receiver]
  relabel_rules    = discovery.relabel.flog_scrape.rules
  refresh_interval = "5s"
}

loki.write "default" {
  endpoint {
    url = "http://loki:3100/loki/api/v1/push"
    tenant_id = "optiop1"
  }
  external_labels = {}
}
```

🟡 **Highlight**: The logs forwarded by this agent are scoped under the **tenant ID: `optiop1`**

---

## 🗃️ Step 3: Create `config2.alloy` for Alloy (System Logs)

```hcl
 local.file_match "local_files" {
     path_targets = [{"__path__" = "/var/log/*.log"}]
     sync_period = "2s"
 }

  loki.source.file "log_scrape" {
    targets    = local.file_match.local_files.targets
    forward_to = [loki.process.filter_logs.receiver]
    tail_from_end = true
  }
  loki.process "filter_logs" {
    stage.drop {
        source = ""
        expression  = ".*Connection closed by authenticating user root"
        drop_counter_reason = "noisy"
      }
    forward_to = [loki.write.grafana_loki.receiver]
    }
  loki.write "grafana_loki" {
    endpoint {
      url = "http://loki:3100/loki/api/v1/push"
      tenant_id = "optiop2"
    }
  }
```

🟡 **Highlight**: This Alloy instance pushes system logs under **tenant ID: `optiop2`**

---

## 🛠️ Step 4: Loki Configuration - `local-config.yaml`

```yaml
auth_enabled: true

server:
  http_listen_port: 3100

common:
  ring:
    instance_addr: 127.0.0.1
    kvstore:
      store: inmemory
  replication_factor: 1
  path_prefix: /tmp/loki

schema_config:
  configs:
  - from: 2020-05-15
    store: tsdb
    object_store: filesystem
    schema: v13
    index:
      prefix: index_
      period: 24h

storage_config:
  filesystem:
    directory: /tmp/loki/chunks
```

🟢 **Important**:  
✅ `auth_enabled: true` is **required** for **multi-tenancy**.  
Each write must specify a `tenant_id`, and queries can be scoped to one.

---
## 🛠️ Step 5: Grafana Configuration - `ds.yaml`
```yaml
apiVersion: 1

datasources:
  - name: optiop1
    type: loki
    access: proxy
    url: http://loki:3100
    jsonData:
      httpHeaderName1: "X-Scope-OrgID"
    secureJsonData:
      httpHeaderValue1: "optiop1"
    isDefault: false
    editable: false

  - name: optiop2
    type: loki
    access: proxy
    url: http://loki:3100
    jsonData:
      httpHeaderName1: "X-Scope-OrgID"
    secureJsonData:
      httpHeaderValue1: "optiop2"
    isDefault: false
    editable: false

```

**In each data source must specify a `jsonData` and `secureJsonData`, for header orgId.**

---
## ▶️ Step 6: Run the Stack

```bash
docker compose up -d
```

---

## 🌐 Access Grafana

Open your browser and go to [http://localhost:3000](http://localhost:3000)

- No login required (anonymous mode is enabled).
- Datasources already provisioned for Loki.

To explore tenants:
- Use `Explore` tab
- Select **Optiop1** data source or **Optiop2**
- Start query logs using builder or code

---

## 🧪 Test It Out

- Start a Docker container with logs and check Alloy-1's behavior.
- Generate system logs (`echo "hello log" >> /var/log/test.log`) and watch Alloy-2 push them to Loki.

---

## ✅ Conclusion

You've successfully set up **multi-tenant logging with Loki and Alloy**!  
