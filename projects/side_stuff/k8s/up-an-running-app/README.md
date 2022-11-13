# The k8s up and running final example project


## Introduction

Just running through the last chapter in the k8s book to get a refresher on kubernetes topics and how to deploy an application. Hoping to use this to help containerize the vodafone project eventually. Going to move over to operators once I finish this up.

## Jupyter

Jupyter notebook is a web-based interactive scientific notebook. All of these projects have their own namespace for organization.

Once the pod is up and running, you need the login token and port forward the application:
```
pod_name=$(kubectl get pods -n jupyter --no-headers | awk '{print $1}')
kubectl logs -n jupyter ${pod_name}
kubectl port-forward -n jupyter ${pod_name} 8888:8888
```
Then, in order to access the application, run `http://localhost:8888/?token=<token>`.