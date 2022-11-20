# Instructions for starting a k8s dashboard

Can follow examples in [this k8s dashboard README](https://github.com/kubernetes/dashboard).

First, create the k8s resources:
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
```

Then, create a proxy for the k8s api server in the bg:
```
kubectl proxy &
```

Open up the dashboard at http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

You will need a bearer token that can be accessed by running
```
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | awk '/default-token/ {print $1}')
```

aaaannnnndddd you should be in!