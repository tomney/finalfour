# AngularFinalfour

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 6.2.3.

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The app will automatically reload if you change any of the source files.

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory. Use the `--prod` flag for a production build.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via [Protractor](http://www.protractortest.org/).

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI README](https://github.com/angular/angular-cli/blob/master/README.md).

## Deploy with Google Kubernetes Engine
You should be able to build the container using 
```
docker build -t gcr.io/bolg-229922/angular-finalfour .
```

To run locally use:
```
docker run -d -p 8080:80 gcr.io/bolg-229922/angular-finalfour 
```

Push this docker image to the container registry via
```
gcloud docker -- push gcr.io/bolg-229922/angular-finalfour
```

Use kubectl to deploy resources to the cluster for the frontend:
```
kubectl apply -f kube-angular.yaml
```

Once deployed the status of the pods can be monitored via:
```
kubectl get pods
```

The status of the service can be monitored via:
```
kubectl describe service angular-finalfour
```

To allow traffic to the nodes, run 