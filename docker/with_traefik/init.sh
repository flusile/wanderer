#!/bin/bash

if [ ! -f .env ]
then
    echo "initialising .env with new generated keys for MEILI and POCKETBASE"
    cp .env.template .env
    echo "MEILI_MASTER_KEY=$(openssl rand -hex 16)" >>.env
    echo "POCKETBASE_ENCRYPTION_KEY=$(openssl rand -hex 16)" >>.env
fi

source .env

mkdir -p ${TRAEFIK_BASE}/tmp
touch ${TRAEFIK_BASE}/acme.json
chmod u=rw,go= ${TRAEFIK_BASE}/acme.json
cp traefik.yaml ${TRAEFIK_BASE}/

mkdir -p ${WANDERER_BASE}/data.ms
mkdir -p ${WANDERER_BASE}/pb_data
mkdir -p ${WANDERER_BASE}/uploads

echo "please change your DOMAIN in the .env file before starting docker compose"
