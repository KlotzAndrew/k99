#!/bin/bash

set -ex

GCP_PROJECT=k99-project
ACCOUNT_NAME=terraform-demo

gcloud iam service-accounts create "$ACCOUNT_NAME" \
  --display-name "Terraform admin account" \
  --project=${GCP_PROJECT}

gcloud iam service-accounts keys create ./terraform-admin.json \
  --iam-account="${ACCOUNT_NAME}"@${GCP_PROJECT}.iam.gserviceaccount.com \
  --project=${GCP_PROJECT}

sudo mv terraform-admin.json ~/.config/gcloud/

gcloud projects add-iam-policy-binding ${GCP_PROJECT} \
  --member serviceAccount:"${ACCOUNT_NAME}"@${GCP_PROJECT}.iam.gserviceaccount.com \
  --role roles/viewer

gcloud projects add-iam-policy-binding ${GCP_PROJECT} \
  --member serviceAccount:"${ACCOUNT_NAME}"@${GCP_PROJECT}.iam.gserviceaccount.com \
 --role roles/storage.admin

gcloud projects add-iam-policy-binding ${GCP_PROJECT} \
  --member serviceAccount:"${ACCOUNT_NAME}"@${GCP_PROJECT}.iam.gserviceaccount.com \
  --role roles/container.admin

gcloud projects add-iam-policy-binding ${GCP_PROJECT} \
  --member serviceAccount:"${ACCOUNT_NAME}"@${GCP_PROJECT}.iam.gserviceaccount.com \
  --role roles/iam.serviceAccountUser
