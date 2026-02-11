# wanderer using docker compose and traefik

Using traefik as a proxy has two benefits:

- it takes care for the https certificates
- it enables the host to server more services than just wanderer

The setup given here uses:

- [traefik](https://traefik.io/traefik) for the routing
- [docker-proxy](https://github.com/Tecnativa/docker-socket-proxy) to have an additional layer of security between traefik and docker
- [watchtower](https://github.com/nicholas-fedor/watchtower) to automatical update images
- and of course the wanderer services

If you do not want to use watchtower you can simply remove that service and any watchtower related label.

The routing in traefik is based on the hostname of the service.
Assume that your host has the hostname myservices.net. The wanderer service then will apperar as wanderer.myservices.net.
Other services, e.g. nextcloud, can be reachable under nextcloud.myservices.net.

For that your dns entry should route any subdomain of your hostname to yor host. Please check taht with you hosting provider.
And beeing there, please check the firewall settings. We need the tcp ports 80, 443 and 8090 open for both ipv4 and ipv6.

Please note: any script and config here is given with best effort. Please make sure you understand what they do.

## initial setup

I assume you are using a kind of linux for your server and your local firewall is acceptinmg the tcp ports mentioned above as well.
And of course we need docker, docker compose and bash.

After copying this folder to your host you can use the init.sh script to setup your environment.
It will create some folders needed and the initial .env file from the .env.template

After this you should check:

- set correct DOMAIN in your .env file. In the example above it was myservices.net.
- set your email adress in traefik/traefik.yaml for letsencrypt

Now you are ready to start the instance using the ./start script given.

To shut it down simply use ./stopp.

After starting it for the first time you have to create a superuser in pocketbase. You can use the following command:

```sh
docker exec wanderer-db /pocketbase superuser upsert **email** **password**
```

## the files in detail

### compose.yaml

Additional to the docker-compose.yaml-Files given in the folder above we add the following:

- internal networks to seperate traffic
- the services mentioned above (traefik, docker-proxy, watchtower)
- we use the internal networks to make sure that meilisearch (wanderer-search) is accessible only by wanderer-web and wanderer-db

### traefik.template.yaml

This file contains the base config for traefik. It will be copied by init.sh as traefik.yaml to the ./traefik folder.

### init.sh

Prepares the environment needed.

Traefik needs a proper config (traefik.yaml) and a safe place for the https certificates.

Wanderer needs some folders to store data in and additiomnal two secret keys for the communication with meilisearch and pocketbase.

You can always every step in here by hand. This script is just for convenience.

### start and stopp

Helper to start and stopp docker compose in a reliable way.

### .env.template

This template contains the environment for docker compose. Some vars are added using init.sh and one possible using the start script.

#### TRAEFIK_BASE

The folder where trafik stores it's config and certificates.

If you want to have it somwhere else on your host, please change the value in the template file.

#### WANDERER_BASE

The folder where the wanderer services store their data.

If you want to have it somwhere else on your host, please change the value in the template file.

#### DOMAIN

The domain of your host. See above for explanation.

#### WANDERER_VER

The image tag for the wanderer docker images.

You can change it to pin your installation to a specific version.

#### PROJECT

Defaults to the name of your directory. Is used by docker compose and traefik to determine networks etc.

You can change it after running init.sh and before using start.

#### Vars defined by wanderer itself

As documented in the general wanderer documentaion there are:

- UPLOAD_USER
- UPLOAD_PASSWORD
- MEILI_MASTER_KEY
- POCKETBASE_ENCRYPTION_KEY

The UPLOAD_*-vars are needed by the cron job for importing track files.
The have to be set with the proper credentials and can be changed at any time.
Changes take effect after starting docker compose the next time

The *_KEY-vars are not part of the template. They will be created and added to the final .env file by the init.sh script.