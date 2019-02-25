If a kubernetes cluster does not yet exist one can be created via
```
gcloud container clusters create bolg \
    --scopes "cloud-platform" \
    --machine-type=g1-small \
    --num-nodes 1 \
    --enable-basic-auth \
    --issue-client-certificate \
    --enable-ip-alias \
    --zone northamerica-northeast1-b
```

Build the docker image via:
```
docker build -t gcr.io/bolg-229922/finalfour .
```

Push this docker image to the container registry via
```
gcloud docker -- push gcr.io/bolg-229922/finalfour
```

Use kubectl to deploy resources to the cluster for the frontend:
```
kubectl create -f deployment/frontend.yaml
```

Use kubectl to deploy resources to the cluster for the load balancer:
```
kubectl create -f deployment/service.yaml
```

Once deployed the status of the pods can be monitored via:
```
kubectl get pods
```

The status of the service can be monitored via:
```
kubectl describe service finalfour-frontend
```
