#!/bin/bash
set -e
# Resolve the container's actual bridge IP so RMI stub advertises the right address.
# hostname -i returns the container IP on the bridge network, not 127.0.0.1.
CONTAINER_IP=$(hostname -i | awk '{print $1}')
exec java \
  -Dcom.sun.management.jmxremote \
  -Dcom.sun.management.jmxremote.port=9999 \
  -Dcom.sun.management.jmxremote.rmi.port=9999 \
  -Dcom.sun.management.jmxremote.authenticate=false \
  -Dcom.sun.management.jmxremote.ssl=false \
  -Dcom.sun.management.jmxremote.local.only=false \
  -Djmx.rmi.registry.port=1099 \
  -Djmx.rmi.port=1099 \
  "-Djava.rmi.server.hostname=${CONTAINER_IP}" \
  -jar /app/app.jar