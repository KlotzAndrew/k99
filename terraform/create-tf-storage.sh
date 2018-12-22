
#!/bin/bash

GCP_PROJECT=k99-project

gsutil mb -p ${GCP_PROJECT} -l us-east1 gs://${GCP_PROJECT}

gsutil versioning set on gs://${GCP_PROJECT}
