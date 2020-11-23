#!/bin/bash -e

ROOT=$(dirname "${BASH_SOURCE[0]}")

LC_CTYPE=C

if [ ! -f "${ROOT}/scripts/env.sh" ]; then
  echo "Please make env.sh file in ./scripts dir from env_example.sh";
fi

echo "Check kubectl and gcloud"

command -v gcloud >/dev/null 2>&1 || \
  { echo >&2 "I require gcloud but it's not installed.  Aborting."; exit 1; }
command -v kubectl >/dev/null 2>&1 || \
  { echo >&2 "I require kubectl but it's not installed.  Aborting."; exit 1; }

echo "Setting Variables"
source "$ROOT"/scripts/env.sh

if [ -z "$PROJECT" ]; then
  echo "Check Your PROJECT variable"
  exit 1
fi

if [ -z "$APP_NAME" ]; then
  echo "Check Your APP_NAME variable"
  exit 1
fi

if [ -z "$APP_VERSION" ]; then
  echo "Check Your APP_VERSION variable"
  exit 1
fi

if [ -z "$CONTAINER_IMAGE" ]; then
  echo "Check Your PROJECT CONTAINER_IMAGE and does it exit"
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
gcloud builds submit --tag "$CONTAINER_IMAGE" .