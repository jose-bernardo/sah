#!/bin/bash

# Check that the CA name has been passed as an argument
#if [ $# -ne 1 ]; then
#    echo "Usage: create-ca <ca-name>"
#    exit 1
#fi

ca_name=CA_sah
ca_dir="./$ca_name"

mkdir "$ca_dir"
cd "$ca_dir"

openssl genrsa -out ca-key.pem 4096

openssl req -x509 -new -nodes -key ca-key.pem -sha256 -days 3600 -out ca.pem -subj "/CN=$ca_name" -addext "subjectAltName=DNS:$ca_name"
