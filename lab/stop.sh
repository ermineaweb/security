#!/bin/bash

docker rm $(docker ps --filter="name=lab-*" -aq) --force