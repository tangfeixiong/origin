#!/bin/bash -e

ORIGIN_ROOT=$(dirname "${BASH_SOURCE}")

if [[ ! -d "$ORIGIN_ROOT/_output/local/bin/linux/amd64/" ]]; then
  echo "please make binaries first"
  exit 1
fi

tgt=$(mktemp -d)

cp -rf $ORIGIN_ROOT/rootfs/* $ORIGIN_ROOT/_output/local/bin/linux/amd64/openshift $ORIGIN_ROOT/_output/local/bin/linux/amd64/oc $tgt

docker build -t tangfeixiong/openshift-origin $tgt

rm -rf $tgt
