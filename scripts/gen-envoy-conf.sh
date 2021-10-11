#!/bin/sh

set -e

echo "Generating envoy.yaml config file..."
cat tmpl/envoy.yaml.tmpl | envsubst \$GATEWAY_ADDRESS > envoy.yaml
