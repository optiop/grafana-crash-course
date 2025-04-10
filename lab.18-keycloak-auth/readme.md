cd /opt/keycloak/bin

./kcadm.sh config credentials --server http://localhost:8080 --realm master --user admin

./kcadm.sh update realms/imbue-sustainability -s sslRequired=NONE --server http://localhost:8080
