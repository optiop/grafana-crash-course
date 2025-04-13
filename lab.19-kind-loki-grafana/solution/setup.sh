#!/bin/bash
# Kill running process 
lsof -ti:3000 | xargs kill -9 || echo "No process found on port 3000"
lsof -ti:12345 | xargs kill -9 || echo "No process found on port 12345"
lsof -ti:3100 | xargs kill -9 || echo "No process found on port 12345"

# run kind cluster
CLUSTER_NAME=$(basename "$(dirname "$(pwd)")")
kind delete cluster --name $CLUSTER_NAME
kind create cluster --name $CLUSTER_NAME

# Add Helm repo
helm repo add grafana https://grafana.github.io/helm-charts

# Install helm charts
helm upgrade --install grafana --namespace grafana --create-namespace grafana/grafana -f ./grafana/values.yaml
helm upgrade --install alloy --namespace alloy --create-namespace grafana/alloy -f ./alloy/values.yaml
helm upgrade --install loki --namespace loki --create-namespace grafana/loki -f ./loki/values.yaml

# Port-forwarding
sleep 10
kubectl port-forward -n grafana service/grafana 3000:80 2>&1 > /dev/null &
kubectl port-forward -n alloy service/alloy 12345:12345 2>&1 > /dev/null &
kubectl port-forward -n loki service/loki-write 3100:3100 2>&1 > /dev/null &

# Retrieve Grafana admin password
GRAFANA_SECRET=$(kubectl get secrets -n grafana grafana -o jsonpath={'.data.admin-password'} | base64 -d)

echo ""
echo "| Grafana | http://localhost:3000 | admin/${GRAFANA_SECRET}"
echo "| Alloy   | http://localhost:12345 | -"
echo "| Loki    | http://localhost:3100 | -"
echo ""