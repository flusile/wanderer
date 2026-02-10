#!/bin/bash

source .env

mkdir ${TRAEFIK_BASE}/tmp
touch ${TRAEFIK_BASE}/acme.json
chmod u=rw,go= ${TRAEFIK_BASE}/acme.json

mkdir ${WANDERER_BASE}/data.ms
mkdir ${WANDERER_BASE}/pb_data
mkdir ${WANDERER_BASE}/uploads

