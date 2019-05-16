#!/bin/bash

set -o -e pipefail

REGISTRY="registry.skycoin.net"

# Build cx-tracker image
if [[ $TRAVIS_BRANCH == "master" ]]; then
  docker build -t $REGISTRY/cx-tracker .
elif [[ $TRAVIS_BRANCH == "develop" ]]; then
  docker build -t $REGISTRY/cx-tracker:develop .
fi
