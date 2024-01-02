#!/usr/bin/env bash

set -euo pipefail

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:latest generate \
    --config /local/.openapi-generator/config.yaml \
    -p enumClassPrefix=true \
    --input-spec /local/api/openapi.yaml \
    --inline-schema-options RESOLVE_INLINE_ENUMS=true \
    --output /local/netbox \
    --http-user-agent go-netbox/$(cat api/netbox_version)
    
