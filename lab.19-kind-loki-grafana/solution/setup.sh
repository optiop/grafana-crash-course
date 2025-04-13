lsof -ti:3000 | xargs kill -9 || echo "No process found on port 3000"

helm repo add grafana https://grafana.github.io/helm-charts

# Install required helm charts
helm upgrade --install grafana --namespace grafana --create-namespace grafana/grafana -f ./grafana/values.yaml
helm upgrade --install alloy --namespace alloy --create-namespace grafana/alloy -f ./alloy/values.yaml



kubectl port-forward -n grafana service/grafana 3000:80 2>&1 > /dev/null &
GRAFANA_SECRET=`kubectl get secrets -n grafana grafana -o jsonpath={'.data.admin-password'} | base64 -d`


kubectl port-forward -n alloy service/alloy 12345:12345 2>&1 > /dev/null &

echo ""
echo ""
echo "| Grafana | http://localhost:3000 | admin/${GRAFANA_SECRET}"
echo "| Alloy   | http://localhost:12345 | -"

