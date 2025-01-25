# Check if a URL is reachable via curl
check_reachable_via_curl() {
  local url=$1
  local regex=$2

  local retry=3
  local delay=4
  echo "  [INFO] Attempting to reach URL: $url and match regex: $regex"
  for i in $(seq 1 $retry); do \
    curl -s $url | grep -q "$regex" \
      && echo "  [PASSED] URL: $url is reachable and matches the regex: $regex" \
      && return 0; \
    sleep $delay; \
  done 
  echo "  [FAILED] URL: $url is not reachable or does not match the regex: $regex"
  return 1
}