#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Usage: <doctor-id>"
    exit 1
fi

doctor_id=$1

mkdir "doctor-$doctor_id"
cd "doctor-$doctor_id"

# Generate the key pair
openssl req -new -newkey rsa:4096 -nodes -keyout doctor-keypair.pem -out doctor-requsest.pem -subj "/CN=doctor-$doctor_id"

# Extract the private key
openssl rsa -in doctor-keypair.pem -out doctor-private.pem

cd ..
# Signing
./sign-certificates.sh "./doctor-$doctor_id/doctor-requsest.pem" "./doctor-$doctor_id/doctor-public.pem"