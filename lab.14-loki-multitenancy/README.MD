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
├── README.md
└── solution
    ├── Alloy/
    │   └── config.alloy
    │   └── config2.alloy
    ├── docker-compose.yaml
    ├── Grafana/
    │   └── provisioning
    │       └── datasources
    │           └── ds.yaml
    └── Loki/
        └── local-config.yaml
```
🟢 **Important**:  
✅ `auth_enabled: true` in Loki config is **required** for **multi-tenancy**.  
Each write must specify a `tenant_id`, and queries can be scoped to one.
**In each Grafana data source  must specify a `jsonData` and `secureJsonData`, for header orgId.**

---
## ▶️ Run the Stack

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
