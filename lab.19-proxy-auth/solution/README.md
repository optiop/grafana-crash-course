```markdown
# Grafana Proxy Auth Example

This project demonstrates how to set up **Grafana with authentication proxy** using **Apache HTTP Server as a reverse proxy** and **basic auth**.

📌 Useful if you want to authenticate users at the proxy layer and pass them to Grafana without exposing Grafana's internal login system.

---

## 🏗️ Project Structure

'''
.
├── apache
│   └── httpd.conf
├── grafana
│   └── grafana.ini
├── .gitignore
├── compose.yml
├── htpasswd
├── LICENSE
└── README.md

'''

---

## ⚙️ How to Use

### 1. Create Basic Auth Credentials

Use `htpasswd` to create a user for Apache Basic Auth:

```bash
htpasswd -bc htpasswd your_username your_password
```

> This creates the `htpasswd` file required by Apache to authenticate users.

---

### 2. Start the Stack

Make sure you have Docker and Docker Compose installed, then run:

```bash
docker-compose up 
```

This spins up:
- **Apache** on `localhost:8080`
- **Grafana** behind the proxy on `localhost:3000` (not directly accessible)

---

### 3. Access the Setup

Go to: [http://localhost:8080](http://localhost:8080)

- Enter the username and password you created.
- You should be redirected to Grafana.
- Grafana will auto-create a user using the username passed by the proxy, based on `auth.proxy` config in `grafana.ini`.

---

---

## 📚 Reference

- 📖 [Grafana Auth Proxy Docs](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-authentication/auth-proxy/)

---

## ✅ Expected Outcome

Once authenticated through Apache:
- The `X-WEBAUTH-USER` header is forwarded to Grafana.
- Grafana automatically logs in or creates the user.
- No internal Grafana login is required.

---
