#!/bin/bash -e

kubectl create namespace "$NAMESPACE"

kubectl create secret --namespace "$NAMESPACE" generic "$DB_SECRET_NAME" \
  --from-literal=username="$DB_USER_NAME" \
  --from-literal=password="$DB_PASSWORD" \
  --from-literal=database="$DB_NAME" \
  --from-literal=port="$DB_PORT" \
  --from-literal=host="$DB_HOST"

echo "Make Credential file in Cluster"

echo "..Create GSA in Google CLoud"
gcloud iam service-accounts create "$GSA_NAME"

echo "..Binding GSA to CloudSQL Client "
gcloud projects add-iam-policy-binding "$PROJECT" \
--member serviceAccount:"$FULL_GSA_NAME" \
--role roles/cloudsql.client

echo "..Generating GSA Credential file"
gcloud iam service-accounts keys create credential.json \
  --iam-account "$FULL_GSA_NAME"

echo "..Create GSA Secret from Credential file"
kubectl create secret --namespace "$NAMESPACE" generic "$GSA_SECRET_NAME" \
--from-file=credential.json=credential.json
