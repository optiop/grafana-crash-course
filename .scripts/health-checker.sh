#!/bin/bash

sleep 30

# Get list of all containers and their status
running_containers=$(docker ps --format "{{.ID}} {{.Names}} {{.State}}")

if [ -z "$running_containers" ]; then
    echo "❌ No running containers found."
    exit 1
fi

echo "🔍 Checking container statuses..."

# Flag to track failures
failed_containers=0

# Iterate over each container and check its state
while IFS= read -r container_info; do
    container_id=$(echo "$container_info" | awk '{print $1}')
    container_name=$(echo "$container_info" | awk '{print $2}')
    container_state=$(docker inspect --format='{{.State.Status}}' "$container_id" 2>/dev/null)

    if [ "$container_state" == "running" ]; then
        echo "✅ Container $container_name ($container_id) is running."
    else
        echo "❌ Container $container_name ($container_id) is in state: $container_state. Failing check!"
        failed_containers=1
    fi
done <<< "$running_containers"

# Exit with failure if any container is not running
if [ "$failed_containers" -eq 1 ]; then
    exit 1
fi

echo "✅ All containers are running successfully."
