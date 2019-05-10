#!/bin/bash

set -o -e pipefail

# Login the private registry
docker login -u "$DOCKER_USERNAME" -p "$DOCKER_PASSWORD" registry.skycoin.net

REGISTRY="registry.skycoin.net"

# Push the image
docker push $REGISTRY/cx-tracker
