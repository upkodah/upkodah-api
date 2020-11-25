#!/bin/bash -e

# Cluster constants
export CLUSTER_NAME="YOUR_CLUSTER_NAME"
export PROJECT="YOUR_PROJECT"
export CLUSTER_ZONE=$(gcloud config get-value compute/zone)
export NAMESPACE="YOUR_APP_NAMESPACE"

# deploy constants
export APP_VERSION="YOUR_APP_VERSION"
export APP_NAME="YOUR_APP_NAME"
export HTTP_PORT="YOUR_HTTP_PORT"
export DEPLOY_NAME="$APP_NAME"-deployment
export CONTAINER_IMAGE=gcr.io/"$PROJECT"/"$APP_NAME":"$APP_VERSION"

# DB constants
export INSTANCE_NAME="YOUR_INSTANCE_NAME"
export DB_NAME="YOUR_DB_NAME"
export DB_USER_NAME="YOUR_DB_USER_NAME"
export DB_PASSWORD="YOUR_DB_PASSWORD"
export DB_HOST="YOUR_DB_HOST"
export DB_PORT="YOUR_DB_PORT"
export DB_SECRET_NAME="$APP_NAME"-cloudsql-creds
export CONNECTION_NAME=$(gcloud sql instances describe "$INSTANCE_NAME" \
--format="value(connectionName)")

# service accounts
export KSA_NAME="$APP_NAME"-ksa
export GSA_NAME="$APP_NAME"-gsa
export FULL_KSA_NAME="$PROJECT".svc.id.goog["$NAMESPACE"/"$KSA_NAME"]
export FULL_GSA_NAME="$GSA_NAME"@"$PROJECT".iam.gserviceaccount.com
