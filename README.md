This project is a simple template for an angular web-app with a Go server. It is currently hardcoded to work for the GKE instance on the Google Cloud Project bolg-229922.

To deploy this webapp to that project clone the repo. Then run the `deploy.sh` script. You will be able to get the external IP of the service by running
```
$ kubectl get services
```

The frontend LoadBalancer service will direct you to the app.