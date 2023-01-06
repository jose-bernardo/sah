#!/bin/bash


mkdir -p "bd-server"
cd "bd-server"

# Generate the key pair
openssl req -new -newkey rsa:4096 -nodes -keyout bd-server-key.pem -out bd-server-request.pem -subj "/CN=bd-server"

# Extract the private key
openssl rsa -in bd-server-key.pem -out bd-server-key.pem

cd ..
openssl x509 -req -sha256 -in ./bd-server/bd-server-request.pem -days 3600 -CA ./CA_sah/ca.pem -CAkey ./CA_sah/ca-key.pem -set_serial 01 -out ./bd-server/bd-server-cert.pem -extfile bd-config.cnf -extensions v3_ca

rm ./bd-server/bd-server-request.pem
cp ./CA_sah/ca.pem ./bd-server/

echo "Signed bd server key"

