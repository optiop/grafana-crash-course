# Build stage
FROM eclipse-temurin:21-jdk AS build

WORKDIR /app

COPY app.java .

RUN javac app.java

# Runtime stage
FROM eclipse-temurin:21-jre

WORKDIR /app

COPY --from=build /app /app

# Download OTel Java Agent
ARG OTEL_JAVA_AGENT_VERSION=2.18.1

ADD https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/download/v${OTEL_JAVA_AGENT_VERSION}/opentelemetry-javaagent.jar /app/opentelemetry-javaagent.jar

ENV JAVA_TOOL_OPTIONS="-javaagent:/app/opentelemetry-javaagent.jar"

EXPOSE 8081

CMD ["java", "app"]