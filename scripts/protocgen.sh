#!/usr/bin/env bash

# How to run manually:
# docker build --pull --rm -f "contrib/devtools/Dockerfile" -t cosmossdk-proto:latest "contrib/devtools"
# docker run --rm -v $(pwd):/workspace --workdir /workspace cosmossdk-proto sh ./scripts/protocgen.sh

set -e

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find ./stratos -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep "option go_package" $file &> /dev/null ; then
    buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done

cd ..
# move proto files to the right places
cp -r github.com/stratosnet/stratos-chain/* ./
rm -rf github.com

./scripts/protocgen-pulsar.sh