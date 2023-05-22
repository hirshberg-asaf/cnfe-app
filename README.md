# CNFE APP

## Instructions

###### 1. Building container images

Under `/fronted` and `/jsonServer` directories, use `docker build . -t <tag-name>` to build each of the microservices
and push them the a container registry. The deployment yaml for both microservices can be found under `/kubernetes` 
and will need to be modified with the container image names.

###### 2. Deploying to Kubernetes

The application consist of two microservices `frontend` and `backend`, a configmap `backend-url` the populate an 
environment variable with the `backend` url in the `frontend`.
This example use `Contour` to expose the `frontend` with `httpproxy` configured in `frontend-httpproxy.yaml`.

First, create the namespace for the application:
```
kubectl apply -f /kubernetes/app-namespace.yaml
``` 

Create the configmap, if you changed the name of the backedn make sure to apply the changes to the value of `url`:
```
kubectl apply -f frontend-configmap.yaml
``` 

Deploy both the `frontend` and `backend` deployments, make sure to change the container image names in both files:
```
kubectl apply -f frontend-deployment.yaml
kubectl apply -f backend-deployment.yaml
```

create a `service` for both the `frontend` and `backend`:
```
kubectl apply -f frontend-service.yaml
kubectl apply -f backend-service.yaml
```

before deploying the `httpproxy` object, ensure to change `fqdn:` to a valid owned `fqdn`.
```
kubectl apply -f frontend-httpproxy.yaml 
```
# Jenkins-on-Kubernetes-guide
