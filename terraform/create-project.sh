#!/bin/bash

GCP_PROJECT=k99-project
# org_id=$(gcloud organizations list)
# billing_acount=$(gcloud beta billing accounts list)

exit 1

gcloud projects create ${GCP_PROJECT} \
  --organization ${org_id} \
  --set-as-default

gcloud beta billing projects link ${GCP_PROJECT} \
  --billing-account ${billing_acount}
