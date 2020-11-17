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

echo "Enable CloudBuild api"
gcloud services enable cloudbuild.googleapis.com

echo "Build Docker File and Submit to ${CONTAINER_IMAGE}"
cd "$ROOT"/../
gcloud builds submit --tag "$CONTAINER_IMAGE" .

cd "$ROOT"/