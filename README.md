# Pod Chrono Scaler

This is a simple program that scales an nginx deployment so that the replicas match the current time's minutes mod 10.

### How to deploy

Assuming you have a deployment in the `default` namespace with `metadata.labels.app` set to `nginx`, simply clone this repo and do the following.

WARNING: This grants blanket edit permissions to the default service account. DO NOT DO THIS IN A PRODUCTION CLUSTER.

```
$ kubectl apply -f clusterrolebinding.yaml
$ kubectl apply -f pod.yaml
```

### How to build

Assuming you have Go 1.14 installed:

```
$ go build -o scaler .
```

We're using an in-cluster config here, so running it locally won't work. But if you want to try anyway:

```
$ go run .
```
