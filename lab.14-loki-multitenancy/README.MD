# 🚀 Multi-Tenant Loki Setup with Grafana Alloy

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
```

---

## 🐳 Step 1: Create `docker-compose.yml`

```yaml
version: "3.9"

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
    image: grafana/alloy:v1.7.5
    privileged: true
    ports:
      - 12346:12345
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:rw
      - ./config2.alloy:/etc/alloy/config.alloy
    networks:
      - app_net
    command:
      - run
      - /etc/alloy/config.alloy
      - --storage.path=/var/lib/alloy/data
      - --server.http.listen-addr=0.0.0.0:12345

  grafana:
    image: grafana/grafana:11.0.0
    ports:
      - "3000:3000"
    environment:
      - GF_PATHS_PROVISIONING=/etc/grafana/provisioning
      - GF_AUTH_ANONYMOUS_ENABLED=true
      - GF_AUTH_ANONYMOUS_ORG_ROLE=Admin
    networks:
      - app_net
    entrypoint:
      - sh
      - -euc
      - |
        mkdir -p /etc/grafana/provisioning/datasources
        cat <<EOF > /etc/grafana/provisioning/datasources/ds.yaml
        apiVersion: 1
        datasources:
        - name: Optiop-1
          type: loki
          access: proxy
          orgId: optiop1
          url: http://loki:3100
          basicAuth: false
          isDefault: false
          version: 1
          editable: false
        - name: Optiop-2
          type: loki
          access: proxy
          orgId: optiop2
          url: http://loki:3100
          basicAuth: false
          isDefault: false
          version: 1
          editable: false
        EOF
        /run.sh

networks:
  app_net:
    name: app_net
```

---
**Note the orgId in the grafana config**
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
local.file_match "example" {
  path_targets = [{
    __address__ = "localhost",
    __path__    = "/var/log/*.log",
  }]
}

loki.source.file "example" {
  targets    = local.file_match.example.targets
  forward_to = [loki.write.default.receiver]
}

loki.write "default" {
  endpoint {
    url = "http://localhost/loki/api/v1/push"
    tenant_id = "optiop2"
  }
  external_labels = {}
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

## ▶️ Step 5: Run the Stack

```bash
docker compose up -d
```

---

## 🌐 Access Grafana

Open your browser and go to [http://localhost:3000](http://localhost:3000)

- No login required (anonymous mode is enabled).
- Datasources already provisioned for Loki.
- **Note that we gave access to already provisioned data source so end user is not able to change Org-id header**

To explore tenants:
- Use `Explore` tab
- Select **Optiop** data source
- Add `X-Scope-OrgID` in **Query Headers** (e.g., `optiop1` or `optiop2`)

---

## 🧪 Test It Out

- Start a Docker container with logs and check Alloy-1's behavior.
- Generate system logs (`echo "hello log" >> /var/log/test.log`) and watch Alloy-2 push them to Loki.

---

## ✅ Conclusion

You've successfully set up **multi-tenant logging with Loki and Alloy**!  
