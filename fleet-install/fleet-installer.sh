#!/usr/bin/env bash

helm -n fleet-system install --create-namespace --wait fleet-crd \
    https://github.com/rancher/fleet/releases/download/v0.3.7/fleet-crd-0.3.7.tgz
helm -n fleet-system install --create-namespace --wait fleet \
    https://github.com/rancher/fleet/releases/download/v0.3.7/fleet-0.3.7.tgz

