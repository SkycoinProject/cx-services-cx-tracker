#!/bin/bash

set -o -e pipefail

REGISTRY="registry.skycoin.net"

# Build cx-tracker image
if [[ $TRAVIS_BRANCH == "master" ]]; then
  docker build -t $REGISTRY/cx-tracker -f docker/images/cx-tracker/Dockerfile .
  docker build -t $REGISTRY/cx-tracker-web -f docker/images/cx-tracker-web/Dockerfile .
elif [[ $TRAVIS_BRANCH == "develop" ]]; then
  docker build -t $REGISTRY/cx-tracker:develop .
  docker build -t $REGISTRY/cx-tracker-web:develop .
fi
