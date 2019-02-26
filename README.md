# FinalFour
Allow users to select the teams that will appear in the NCAA final four. 

This project is currently hardcoded to work for the GKE instance on the Google Cloud Project bolg-229922.

To deploy this webapp to that project clone the repo. Then run the `deploy.sh` script. You will be able to get the external IP of the service by running
```
$ kubectl get services
```

Enter the external IP for the frontend load balancer into your browser's navbar to go to the app.
