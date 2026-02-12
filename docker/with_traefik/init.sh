#!/bin/bash

if [ ! -f .env ]
then
    echo "initialising .env with new generated keys for MEILI and POCKETBASE"
    cp .env.template .env
    echo "MEILI_MASTER_KEY=$(openssl rand -hex 16)" >>.env
    echo "POCKETBASE_ENCRYPTION_KEY=$(openssl rand -hex 16)" >>.env
fi

source .env

# traefik needs a tmp directory
mkdir -p ${TRAEFIK_BASE}/tmp
# and the acme.json file has to exist with proper rights
touch ${TRAEFIK_BASE}/acme.json
chmod u=rw,go= ${TRAEFIK_BASE}/acme.json

cp traefik.template.yaml ${TRAEFIK_BASE}/traefik.yaml

# wanderer also needs volumes
# wo don't use docker volumes here to be able to put them on physical media of our choice
mkdir -p ${WANDERER_BASE}/data.ms
mkdir -p ${WANDERER_BASE}/pb_data
mkdir -p ${WANDERER_BASE}/uploads

cat <<+++
please change
- your DOMAIN in the .env file
- your email address in traefik/traefik.yaml
before starting docker compose
+++
