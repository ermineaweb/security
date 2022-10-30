#!/bin/bash

# docker run --rm -d \
# -p 3000:3000 \
# --name lab-juice-shop \
# bkimminich/juice-shop

# docker run --rm -d \
# -p 4000:80 \
# --name lab-dvwa \
# vulnerables/web-dvwa

docker run --rm -d \
-p 5000:80 \
--name lab-metasploitable2 \
tleemcjr/metasploitable2

watch docker ps