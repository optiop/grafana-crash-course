#!/usr/bin/env bash
set -euo pipefail

CLUSTER_NAME="java-agent-demo"

if [[ -f .pf-pids ]]; then
  rm -f .pf-pids
fi

echo ">> Delete kind cluster: $CLUSTER_NAME"
kind delete cluster --name "$CLUSTER_NAME" || true