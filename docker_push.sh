#!/bin/bash

set -o -e pipefail

REGISTRY="registry.skycoin.net"

# Login the private registry
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin $REGISTRY

# Push the images
if [[ $TRAVIS_BRANCH == "develop" ]]; then
  docker push $REGISTRY/cx-tracker
  docker push $REGISTRY/cx-tracker-web
fi
