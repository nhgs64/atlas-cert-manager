
# GlobalSign Atlas cert-manager External Issuer

External issuers extend [cert-manager](https://cert-manager.io/) to issue certificates using APIs and services
which aren't built into the cert-manager core.

This repository implements an [External Issuer] for GlobalSign's Atlas .

## Install

```console
kubectl apply -f https://github.com/cert-manager/sample-external-issuer/releases/download/v0.1.0/install.yaml
```



## Links

[External Issuer]: https://cert-manager.io/docs/contributing/external-issuers
[cert-manager Concepts Documentation]: https://cert-manager.io/docs/concepts
[Kubebuilder Book]: https://book.kubebuilder.io
[Kubebuilder Markers]: https://book.kubebuilder.io/reference/markers.html
[Distroless Docker Image]: https://github.com/GoogleContainerTools/distroless
[Configure a Security Context]: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
[kube-rbac-proxy]: https://github.com/brancz/kube-rbac-proxy
[GitHub New Release Page]: https://github.com/cert-manager/sample-external-issuer/releases/new
