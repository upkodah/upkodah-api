#!/bin/bash -e

kubectl create namespace "$NAMESPACE"

kubectl create secret --namespace "$NAMESPACE" generic "$DB_SECRET_NAME" \
  --from-literal=username="$DB_USER_NAME" \
  --from-literal=password="$DB_PASSWORD" \
  --from-literal=database="$DB_NAME" \
  --from-literal=port="$DB_PORT" \
  --from-literal=host="$DB_HOST"

echo "Enable Workload Identity in Cluster"

echo "..Enable Workload"
gcloud container clusters update "$CLUSTER_NAME" \
  --workload-pool="$PROJECT".svc.id.goog

#echo "..Create Node Pool"
#gcloud container node-pools create "$NODEPOOL_NAME" \
#  --cluster="$CLUSTER_NAME" \
#  --workload-metadata=GKE_METADATA

echo "..Update Base Node Pools"
gcloud container node-pools update "$NODEPOOL_NAME" \
  --cluster="$CLUSTER_NAME" \
  --workload-metadata=GKE_METADATA

echo "..Authenticating to Google Cloud"

echo "....Create SA in Kubernetes Cluster"
gcloud container clusters get-credentials "$CLUSTER_NAME"
kubectl create serviceaccount --namespace "$NAMESPACE" "$KSA_NAME"

echo "....Create GSA in Google CLoud"
gcloud iam service-accounts create "$GSA_NAME"

echo "....Binding to Google Cloud"
gcloud iam service-accounts add-iam-policy-binding \
  --role roles/iam.workloadIdentityUser \
  --member "serviceAccount:${FULL_KSA_NAME}" \
  "$FULL_GSA_NAME"

echo "....Binding to Kubernetes"
kubectl annotate serviceaccount \
  --namespace "$NAMESPACE" \
  "$KSA_NAME" \
  iam.gke.io/gcp-service-account="$FULL_GSA_NAME"

#echo "..Test Workload Identity"
#kubectl run -it \
#  --image google/cloud-sdk:slim \
#  --serviceaccount "$KSA_NAME" \
#  --namespace "$NAMESPACE" \
#  workload-identity-test
#
## Then run this code
#gcloud auth list
#exit

