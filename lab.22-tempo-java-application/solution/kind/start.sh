#!/usr/bin/env bash
set -euo pipefail

CLUSTER_NAME="java-agent-demo"

echo ">> Create kind cluster: $CLUSTER_NAME"
kind create cluster --name "$CLUSTER_NAME"

echo ">> Create namespace"
kubectl create ns observability

echo ">> Build & load java app image"
docker build -t java-app:latest ./app
kind load docker-image java-app:latest --name "$CLUSTER_NAME"

echo ">> Apply deployment"
kubectl apply -n observability -f deployment/tempo.yaml
kubectl apply -n observability -f deployment/grafana.yaml
kubectl apply -n observability -f deployment/app.yaml

echo ">> Wait for pods"
kubectl -n observability wait --for=condition=available deploy/tempo --timeout=120s
kubectl -n observability wait --for=condition=available deploy/grafana --timeout=120s
kubectl -n observability wait --for=condition=available deploy/java-app --timeout=120s

echo ">> Port-forward to localhost"
kubectl -n observability port-forward svc/grafana 3000:3000 >/dev/null 2>&1 &
PF1=$!
kubectl -n observability port-forward svc/tempo 3200:3200 >/dev/null 2>&1 &
PF2=$!
kubectl -n observability port-forward svc/java-app 8081:8081 >/dev/null 2>&1 &
PF3=$!

echo $PF1 $PF2 $PF3 > .pf-pids

echo
echo "Grafana   → http://localhost:3000"
echo "Tempo UI  → http://localhost:3200"
echo "Java App  → http://localhost:8081"