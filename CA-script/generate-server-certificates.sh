#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Usage: <server-name>"
    exit 1
fi

server_name=$1

mkdir "server-$server_name"
cd "server-$server_name"

# Generate the key pair
openssl req -new -newkey rsa:4096 -nodes -keyout server-key.pem -out server-request.pem -subj "/CN=server-$server_name"

# Extract the private key
openssl rsa -in server-key.pem -out server-private.pem

cd ..
# Signing
./sign-certificates.sh "./server-$server_name/server-request.pem" "./server-$server_name/server-cert.pem"