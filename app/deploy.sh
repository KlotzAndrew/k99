#!/bin/bash

set -ex

APP_NAME=k99
DEPLOYS=$(helm ls | grep "$APP_NAME" | wc -l)

if [ ${DEPLOYS}  -eq 0 ]; then
  helm install --name "$APP_NAME" --namespace=demo ./chart
else
  helm upgrade "$APP_NAME" --namespace=demo ./chart
fi
