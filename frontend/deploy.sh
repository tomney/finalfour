# Install the node module dependencies
npm install

# Build the angular project
ng build

# Build the docker image via:
docker build -t gcr.io/bolg-229922/angular-finalfour .

# RUNNING LOCALLY CAN BE DONE VIA
# docker run -d -p 8080:80 gcr.io/bolg-229922/angular-finalfour:latest

# Push this docker image to the container registry via
gcloud docker -- push gcr.io/bolg-229922/angular-finalfour

# Use kubectl to deploy resources to the cluster for the backend:
kubectl create -f frontend.yaml