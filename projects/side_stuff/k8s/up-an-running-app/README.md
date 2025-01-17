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

## Ghost

Ghost is a popular blogging engine with a clean interface written in javascript. In order to get the site up and running, we need a ghost config file. We will mount this config file as a config map in our deployment. In order to deploy this config map, run `kubectl create cm --from-file ghost-config.js ghost-config`.

Within the ghost deployment object, we define this config map as our volume mount in the pod template. Once the deployment is up and running, we will expose the deployment with the command `kubectl expose deployments ghost --port=2368`. In order to access the blog, now run `kubectl proxy`.

Couldn't get it working and want to move on. As much as it would be nice to get the troubleshooting skills, it's a goddamn Sunday and I wanna do this redis example and hit the gym.

## Redis

Redis is an in-memory k/v store(obvi). Going to deploy an example one into a kubernetes cluter! In order to deploy this, we will need to deploy a `redis-server` which is the k/v store and `redis-sentinel`, which implements the health checking and failover logic for the replicated redis cluster. A redis cluster uses the parent-child relationship to provide fail-over and replication capabilites. Let's do it!

In order to initialize redis, we need to map all of the config files and the initialization scripts into a configmap:
```
kubectl create configmap \
--from-file=slave.conf=./slave.conf \
--from-file=master.conf=./master.conf \
--from-file=sentinel.conf=./sentinel.conf \
--from-file=init.sh=./init.sh \
--from-file=sentinel.sh=./sentinel.sh \
redis-config
```

The redis service created provides naming and discovery for the redis replicas(in this case, just the one redis server). In the redis object, there are two volumes. One for the configmap to configure the two redis applications and an emptyDir volume that is mapped into the redis server container to hold the application data so that it survives a container restart.

Couldn't get it working, but I should go to the gym if I want to make it home in time to help with dinner.