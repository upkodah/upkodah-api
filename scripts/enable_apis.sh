#!/bin/bash -e

echo "Enabling GoogleApiServices"

function enable_api() {
  SERVICE=$1
  if [[ $(gcloud services list --format="value(serviceConfig.name)" \
                                --filter="serviceConfig.name:$SERVICE" 2>&1) != \
                                "$SERVICE" ]]; then
    echo "Enabling $SERVICE"
    gcloud services enable "$SERVICE"
  else
    echo "$SERVICE is already enabled"
  fi
}

enable_api sqladmin.googleapis.com
enable_api container.googleapis.com
enable_api iam.googleapis.com
