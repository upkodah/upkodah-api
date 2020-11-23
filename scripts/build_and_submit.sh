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

if [ -z "$PROJECT" ]; then
  echo "Check Your PROJECT variable"
  exit 1
fi


CLOUD_BUILD=cloudbuild.googleapis.com
if [[ $(gcloud services list --format="value(serviceConfig.name)" \
                              --filter="serviceConfig.name:${CLOUD_BUILD}" 2>&1) != \
                              "$CLOUD_BUILD" ]]; then
  echo "Enabling $CLOUD_BUILD"
  gcloud services enable "$CLOUD_BUILD"
else
  echo "$CLOUD_BUILD is already enabled"
fi

echo "Build Docker File and Submit to ${CONTAINER_IMAGE}"
gcloud builds submit --tag "$CONTAINER_IMAGE" ..