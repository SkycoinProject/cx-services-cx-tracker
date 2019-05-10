#!/bin/bash

set -o -e pipefail

REGISTRY="registry.skycoin.net"

# Build cx-tracker image
docker build -t $REGISTRY/cx-tracker .
