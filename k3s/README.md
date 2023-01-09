## How to deploy
As prerequisites, you need to have a Kubernetes cluster and an AWS S3 bucket set up. 
In addition, your aws CLI tool has been configured with the right credentials (ACCESS_KEY_ID, SECRET_ACCESS_KEY)

### Kubernetes cluster
Make sure you have a k8s cluster up and running. It could KIND, minikube, or my favorite, k3s. 
The easiest way to deploy and managed k3s clusters is through [k3d](https://k3d.io/).

### Setup
Make sure you have created an S3 bucket and note the bucket name. Populate the bucket name in litestream.yml in the placeholder in the url field.
With the right Kubeconfig pointing to your k8s cluster, run:
* `configure-litestream.sh` to set up the namespace, secret and configmap
* `kubectl apply -f sts.yaml -n sqlitestream `

