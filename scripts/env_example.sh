#!/bin/bash -e

# deploy constants
export APP_VERSION="YOUR_APP_VERSION"
export APP_NAME="YOUR_APP_NAME"
export DEPLOY_NAME="$APP_NAME"-deployment
export CONTAINER_IMAGE="CONTAINER_IMAGE"

# Cluster constants
export NODEPOOL_NAME="YOUR_NODEPOOL_NAME"
export CLUSTER_NAME="YOUR_CLUSTER_NAME"
export PROJECT="YOUR_PROJECT"
export CLUSTER_ZONE=$(gcloud config get-value compute/zone)
export NAMESPACE="YOUR_APP_NAMESPACE"

# DB constants
export INSTANCE_NAME="YOUR_INSTANCE_NAME"
export DB_NAME="YOUR_DB_NAME"
export DB_USER_NAME="YOUR_DB_USER_NAME"
export DB_PASSWORD="YOUR_DB_PASSWORD"
export DB_SECRET_NAME="$APP_NAME"-cloudsql-creds
export DB_PORT=3306
export CONNECTION_NAME=$(gcloud sql instances describe "$INSTANCE_NAME" \
--format="value(connectionName)")

# service accounts
export KSA_NAME="$APP_NAME"-ksa
export GSA_NAME="$APP_NAME"-gsa
export FULL_SA_NAME="$PROJECT".svc.id.goog["$NAMESPACE"/"$SA_NAME"]
export FULL_GSA_NAME="$GSA_NAME"@"$PROJECT".iam.gserviceaccount.com
