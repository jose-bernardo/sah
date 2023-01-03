#!/bin/bash

# Check that a public key file has been passed as an argument
if [ $# -ne 2 ]; then
    echo "Usage: sign-key <public-key-file> <certificate-output-file>"
    exit 1
fi

# Set the key file and signature file names
key_file=$1
output_file=$2

# Check that the key file exists
if [ ! -f "$key_file" ]; then
    echo "Error: Key file $key_file not found"
    exit 1
fi

# Sign the CSR and create a certificate
openssl x509 -req -days 3600 -in $key_file -CA ./CA_sah/ca.pem -CAkey ./CA_sah/ca-key.pem -set_serial 01 -out "$output_file"


echo "Successfully signed key file $key_file"
