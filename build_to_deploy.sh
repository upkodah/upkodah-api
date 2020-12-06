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

echo "Get Cluster Credentials"
gcloud container clusters get-credentials "$CLUSTER_NAME" --zone "$CLUSTER_ZONE"

DEPLOY_TEMPLATE=$(< "$ROOT"/deploy/deploy_template.yaml sed \
-e "s,{{ .DEPLOY_NAME }},${DEPLOY_NAME},g" \
-e "s,{{ .APP_NAME }},${APP_NAME},g" \
-e "s,{{ .CONTAINER_IMAGE }},${CONTAINER_IMAGE},g" \
-e "s,{{ .DB_SECRET_NAME }},${DB_SECRET_NAME},g" \
-e "s,{{ .DB_PORT }},${DB_PORT},g" \
-e "s,{{ .CONNECTION_NAME }},${CONNECTION_NAME},g" \
-e "s,{{ .GSA_SECRET_NAME }},${GSA_SECRET_NAME},g" \
-e "s,{{ .GSA_SECRET_VOLUME }},${GSA_SECRET_VOLUME},g" \
-e "s,{{ .BUS_SERVICE_KEY }},${BUS_SERVICE_KEY},g" \
-e "s,{{ .METRO_SERVICE_KEY }},${METRO_SERVICE_KEY},g"
)
echo "$DEPLOY_TEMPLATE" | kubectl apply -n "$NAMESPACE" -f -