#!/bin/bash -e

ROOT=$(dirname "${BASH_SOURCE[0]}")

LC_CTYPE=C

if [ ! -f './env.sh' ]; then
  echo "Please make setting.sh file from setting_example.sh";
fi

echo "Check kubectl and gcloud"

command -v gcloud >/dev/null 2>&1 || \
  { echo >&2 "I require gcloud but it's not installed.  Aborting."; exit 1; }
command -v kubectl >/dev/null 2>&1 || \
  { echo >&2 "I require kubectl but it's not installed.  Aborting."; exit 1; }

echo "Setting Variables"
source "$ROOT"/env.sh

if [ -z "$CLUSTER_ZONE" ]; then
  echo "Make sure that compute/zone is set in gcloud config"
  exit 1
fi

if [ -z "$CONNECTION_NAME" ]; then
  echo "Check Connection Setting of Your CloudSQL"
  exit 1
fi

echo "Get Cluster Credentials"
gcloud container clusters get-credentials "$CLUSTER_NAME" --zone "$CLUSTER_ZONE"

# Setting Components
sh "$ROOT"/enable_apis.sh
sh "$ROOT"/sql_proxy.sh
