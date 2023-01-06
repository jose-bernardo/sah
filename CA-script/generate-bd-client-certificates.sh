#!/bin/bash


mkdir -p "bd-client"
cd "bd-client"

# Generate the key pair
openssl req -new -newkey rsa:4096 -nodes -keyout bd-client-key.pem -out bd-client-request.pem -subj "/CN=bd-client"

# Extract the private key
openssl rsa -in bd-client-key.pem -out bd-client-key.pem

cd ..
#openssl x509 -req -sha256 -in ./bd-client/bd-client-request.pem -days 3600 -CA ./CA_sah/ca.pem -CAkey ./CA_sah/ca-key.pem -set_serial 01 -out ./bd-client/bd-client-cert.pem -extfile bd-config.cnf -extensions v3_ca

# Signing
./sign-certificates.sh "./bd-client/bd-client-request.pem" "./bd-client/bd-client-cert.pem"

rm ./bd-client/bd-client-request.pem

cp ./CA_sah/ca.pem ./bd-client/
echo "Signed bd client key"

