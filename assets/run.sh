#!/bin/bash

cd deployments/

docker-compose --env-file ../configs/config.env up --build