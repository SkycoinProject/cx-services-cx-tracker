#!/bin/bash

set -o -e pipefail

REGISTRY="registry.skycoin.net"

# Build cx-tracker image
if [[ $TRAVIS_BRANCH == "master" ]]; then
	docker build -t $REGISTRY/cx-tracker .
else
	docker build -t $REGISTRY/cx-tracker:$TRAVIS_BRANCH .
fi
