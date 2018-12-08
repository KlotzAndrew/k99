##### create gke cluster

# (by hand) create project

gcloud container clusters create k99 \
  --cluster-version=1.11.3-gke.18 \
  --zone=us-east1-b \
  --num-nodes=3 \
  --machine-type=n1-highcpu-4 \
  --preemptible \
  --no-enable-cloud-logging \
  --disk-size=30 \
  --enable-autorepair \
  --scopes=gke-default,compute-rw,storage-rw \
  --metadata disable-legacy-endpoints=true \
  --no-enable-basic-auth \
  --no-issue-client-certificate \
  --enable-ip-alias

gcloud config set project k99-project

gcloud container clusters get-credentials k99 -z=us-east1-b

kubectl create clusterrolebinding "cluster-admin-$(whoami)" \
  --clusterrole=cluster-admin \
  --user="$(gcloud config get-value core/account)"

# validate setup
kubectl get nodes -o wide

##### DNS

# ns-109.awsdns-13.com.
# ns-1637.awsdns-12.co.uk.
# ns-1169.awsdns-18.org.
# ns-914.awsdns-50.net.

# (by hand, before) need hostname
gcloud dns managed-zones create \
  --dns-name="mixship.com." \
  --description="Istio zone" "istio"

gcloud compute addresses create istio-gateway-ip --region us-east1
gcloud compute addresses describe istio-gateway-ip --region us-east1

DOMAIN="mixship.com"
GATEWAYIP="35.231.8.197"

gcloud dns record-sets transaction start --zone=istio
gcloud dns record-sets transaction add --zone=istio \
  --name="${DOMAIN}" --ttl=300 --type=A ${GATEWAYIP}

gcloud dns record-sets transaction add --zone=istio \
  --name="www.${DOMAIN}" --ttl=300 --type=A ${GATEWAYIP}

gcloud dns record-sets transaction add --zone=istio \
  --name="*.${DOMAIN}" --ttl=300 --type=A ${GATEWAYIP}
gcloud dns record-sets transaction execute --zone istio

# (by hand) set domain registar to hosted zone

##### Install Istio

curl -L https://git.io/getLatestIstio | sh -
cd istio-1.0.4/
sudo cp ./bin/istioctl /usr/local/bin/istioctl

kubectl apply -f ./install/kubernetes/helm/helm-service-account.yaml

# gcloud in docker has file permission issues???
# sudo chmod +r ~/.config/gcloud/ -R

helm init --service-account tiller

gcloud container clusters describe k99 --zone=us-east1-b \
  | grep -e clusterIpv4Cidr -e servicesIpv4Cidr

# (by hand) update config w/ ip + CIDR blocks
helm install --name istio \
  ../../../istio-1.0.4/install/kubernetes/helm/istio \
  --namespace=istio-system \
  -f istio-config.yaml

kubectl -n istio-system get pods

##### Let's Encrypt!

kubectl apply -f ./istio-gateway.yaml

# create dns-admin service acc w/ key in a secret
GCP_PROJECT=k99-project
gcloud iam service-accounts create dns-admin \
  --display-name=dns-admin \
  --project=${GCP_PROJECT}

gcloud iam service-accounts keys create ./gcp-dns-admin.json \
  --iam-account=dns-admin@${GCP_PROJECT}.iam.gserviceaccount.com \
  --project=${GCP_PROJECT}

gcloud projects add-iam-policy-binding ${GCP_PROJECT} \
  --member=serviceAccount:dns-admin@${GCP_PROJECT}.iam.gserviceaccount.com \
  --role=roles/dns.admin

kubectl create secret generic cert-manager-credentials \
  --from-file=./gcp-dns-admin.json \
  --namespace=istio-system

# setup LE
kubectl apply -f ./letsencrypt-issuer.yaml

kubectl apply -f ./wildcard-cert.yaml

# expose virtual service

kubectl apply -f ./grafana-virtual-service.yaml

curl -I --http2 https://grafana.mixship.com

### A/B testing

helm repo add sp https://stefanprodan.github.io/k8s-podinfo

kubectl apply -f ./demo-namespace.yaml

# setup services

helm install --name frontend sp/podinfo-istio \
  --namespace demo \
  -f ./frontend.yaml

helm install --name backend sp/podinfo-istio \
  --namespace demo \
  -f ./backend.yaml

helm install --name store sp/podinfo-istio \
  --namespace demo \
  -f ./store.yaml

# start deploy

helm upgrade --install frontend sp/podinfo-istio \
  --namespace demo \
  -f ./frontend-green.yaml

helm upgrade --install backend sp/podinfo-istio \
  --namespace demo \
  -f ./backend-green.yaml

helm upgrade --install store sp/podinfo-istio \
  --namespace demo \
  -f ./store.yaml

kubectl apply -f ./demo-rules.yaml
