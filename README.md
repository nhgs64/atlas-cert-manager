
# GlobalSign Atlas cert-manager External Issuer

External issuers extend [cert-manager](https://cert-manager.io/) to issue certificates using APIs and services
which aren't built into the cert-manager core.

This repository implements an [External Issuer] for GlobalSign's Atlas .

## Install

```console

```
## Building
### Prerequisites
You will need the following command line tools installed on your PATH:

* [Git](https://git-scm.com/)
* [Golang v1.17+](https://golang.org/)
* [Docker v17.03+](https://docs.docker.com/install/)
* [Kubectl v1.11.3+](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
* [Kubebuilder v2.3.1+](https://book.kubebuilder.io/quick-start.html#installation)
* [Kustomize v3.8.1+](https://kustomize.io/)

### Build
```
make manager

## Usage

There are sample yaml files in the samples directory. To start issuing, an Atlas issuer needs to be deployed along with a secret.
The secret (see example sampels/secret_issuer.yaml) contains four fields which must contain the base64 encoded API key, API secret, 
mTLS cert and mTLS key.
```
cat mymtlscertificate.pem | base64 -w 0
cat mymtlskey.pem | base64 -w 0
echo 1234567abcdef | base64                        # encode the API key
echo en82u8uXmo39u94uFG9589489djiJdid | base64     # encode the API secret
```
*Note: certificate and key are expected to be in PEM format, not DER*

Next, deploy the secret and issuer:
```
kubectl create -f samples/secret_issuer.yaml
kubectl create -f sample/sample-issuer_v1alpha1_issuer.yaml
```
Kubernetes is now ready to issue Atlas certificates. Certificate and certificate request objects can be created the same way 
as other cert-manager issuers, however the group in the issuerRef must specify `hvca.globalsign.com`. See samples/certificate_issuer.yaml
for an example.

## Links

[External Issuer]: https://cert-manager.io/docs/contributing/external-issuers
[cert-manager Concepts Documentation]: https://cert-manager.io/docs/concepts
[Kubebuilder Book]: https://book.kubebuilder.io
[Kubebuilder Markers]: https://book.kubebuilder.io/reference/markers.html
[Distroless Docker Image]: https://github.com/GoogleContainerTools/distroless
[Configure a Security Context]: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
[kube-rbac-proxy]: https://github.com/brancz/kube-rbac-proxy
[GitHub New Release Page]: https://github.com/cert-manager/sample-external-issuer/releases/new
