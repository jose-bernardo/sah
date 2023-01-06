#!/bin/bash

#if [ $# -ne 1 ]; then
#    echo "Usage: <server-name>"
#    exit 1
#fi

server_name="INTERNAL_SAH_SERVER"

mkdir "$server_name"
cd "$server_name"

# Generate the key pair
openssl req -new -newkey rsa:4096 -nodes -keyout server-key.pem -out server-request.pem -subj "/CN=$server_name"

# Extract the private key
openssl rsa -in server-key.pem -out server-private.pem

cd ..
# Signing
./sign-certificates.sh "./$server_name/server-request.pem" "./$server_name/server-cert.pem"
rm "./$server_name/server-request.pem"
