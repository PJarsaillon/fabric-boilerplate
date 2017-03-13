#!/bin/bash
./clean.sh
docker-compose build
docker-compose up
exit 0