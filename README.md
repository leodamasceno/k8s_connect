# k8s connect
This repository stores the binary and the source code of the Golang application k8s_connect, which is used to connect to the AWS EKS clusters.

## Running it
Clone the repository and execute the commands below:
```
# export KUBECONFIG=$HOME/.kube/config
# mv bin/k8s_connect-0.1 /usr/local/bin/k8s_connect
# k8s_connect dev
```

## Building a new version
After changing the *main.go* file, a new version can be released by following the steps below:
```
# cd src/k8s_connect/
# go build
# mv k8s_connect ../../bin/k8s_connect-NEW_VERSION
```
