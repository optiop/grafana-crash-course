# Java Zero Code Agent

## for run application
``` bash 
cd java/
java \
  -javaagent:./opentelemetry-javaagent.jar \
  -Dotel.service.name=simple-java-app \
  -Dotel.traces.exporter=otlp \
  -Dotel.exporter.otlp.protocol=grpc \
  -Dotel.exporter.otlp.endpoint=http://localhost:4317 \
  -Dotel.exporter.otlp.insecure=true \
  app
  ```