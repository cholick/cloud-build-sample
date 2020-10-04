# cloud-build-sample

This is a sample Cloud Build pipeline that takes Go source and
* Unit tests
* Compiles
* Lints
* Builds and pushes images
* Deploys to a Kubernetes cluster
* Tests that deployed workload

### Pre-reqs
#### Build the builder

```bash
gcloud builds submit deployment/helm-builder/ --config=deployment/helm-builder/cloudbuild.yaml
```

Builder [source](https://github.com/GoogleCloudPlatform/cloud-builders-community/tree/master/helm).
* Customized image name and version in `cloudbuild.yaml`
* Removed Helm 2 stuff from `helm.bash`

Permissions:
https://github.com/GoogleCloudPlatform/cloud-builders/tree/master/kubectl#using-this-builder-with-google-kubernetes-engine

#### Create a GKE Cluster

The `install-dev` step has environment variables that correspond to the cluster

#### Install Ingress

Install the Ingress Helm chart. This code below consumes a reserved static IPv4 address.

```bash
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install my-ingress ingress-nginx/ingress-nginx --set controller.service.loadBalancerIP=XX.XX.XX.XXX
```

### Customize

The `e2e_test.py` and `values-staging.yaml` file use staging.example.com as the domain they're deployed to; actually
using this code would require the value to be change.

### References
* Variable references
    * https://cloud.google.com/cloud-build/docs/configuring-builds/substitute-variable-values
* Full yaml file description
    * https://cloud.google.com/cloud-build/docs/build-config
