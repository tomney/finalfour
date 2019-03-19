
## Deployment
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

To use the Cloud SQL database secrets will need to be established for the username and password. Enter the username of the Cloud SQL user as a string in a file called `username.txt`. Do the same for the password in a file called `password.txt`. These secrets can be created on your kubernetes cluster using:
```
kubectl create secret generic db-user-pass --from-file=./username.txt --from-file=./password.txt
```
These values will be created as environment variables in `backend.yaml`.

We will also need the Cloud Database certificates to be created as a Kubernetes secret. Gather the `client-cert.pem`, `client-key.pem` and `server-ca.pem` files into one folder and run:
```
kubectl create secret generic db-certs --from-file=./client-cert.pem --from-file=./client-key.pem --from-file=./server-ca.pem
```

## Testing
Tests can be run via
```
go test ./...
```

When writing tests mocks can be created using mockery. This will require mockery to be installed in your local GOPATH via:
```
go get github.com/vektra/mockery/.../
```

Then mocks can be created for the appropriate interfaces by navigating to the folder directory then running 
```
$GOPATH/bin/mockery
```