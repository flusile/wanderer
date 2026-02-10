#!/bin/bash

source .env

mkdir -p ${TRAEFIK_BASE}/tmp
touch ${TRAEFIK_BASE}/acme.json
chmod u=rw,go= ${TRAEFIK_BASE}/acme.json
cp traefik.yaml ${TRAEFIK_BASE}/

mkdir -p ${WANDERER_BASE}/data.ms
mkdir -p ${WANDERER_BASE}/pb_data
mkdir -p ${WANDERER_BASE}/uploads

