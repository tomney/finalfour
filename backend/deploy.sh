# Build the docker image via:
docker build -t gcr.io/bolg-229922/finalfour .

# RUNNING LOCALLY CAN BE DONE VIA
# docker run -d -p 8080:80 gcr.io/bolg-229922/finalfour:latest

# Push this docker image to the container registry via
gcloud docker -- push gcr.io/bolg-229922/finalfour

# Use kubectl to deploy resources to the cluster for the backend:
kubectl create -f backend.yaml
kubectl create -f backend-service.yaml
