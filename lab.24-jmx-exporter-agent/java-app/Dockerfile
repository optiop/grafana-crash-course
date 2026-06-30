# ── Stage 1: Build ──────────────────────────────────────────────────────────
FROM maven:3.9-eclipse-temurin-17 AS builder
WORKDIR /app

ARG JMX_AGENT_VERSION=0.20.0

# Cache Maven dependencies first
COPY pom.xml .
RUN mvn dependency:go-offline -q

# Build app
COPY src/ ./src/
RUN mvn package -q -DskipTests

# Download JMX Prometheus agent
RUN curl -fsSL \
    "https://repo1.maven.org/maven2/io/prometheus/jmx/jmx_prometheus_javaagent/${JMX_AGENT_VERSION}/jmx_prometheus_javaagent-${JMX_AGENT_VERSION}.jar" \
    -o target/jmx_prometheus_javaagent.jar

# ── Stage 2: Runtime ─────────────────────────────────────────────────────────
FROM eclipse-temurin:17-jre
WORKDIR /app

COPY --from=builder /app/target/hello-world-jmx-1.0-SNAPSHOT.jar app.jar
COPY --from=builder /app/target/jmx_prometheus_javaagent.jar jmx_prometheus_javaagent.jar
COPY jmx-config.yaml .

# 8080 = Hello World HTTP, 9090 = JMX Exporter Prometheus metrics
EXPOSE 8080 9090

ENTRYPOINT ["java", \
  "-javaagent:/app/jmx_prometheus_javaagent.jar=9090:/app/jmx-config.yaml", \
  "-jar", "/app/app.jar"]
