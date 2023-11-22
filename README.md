### Overview
Cronjob CR and controller that schedules native batchv1/cronjobs.
The API supports two versions of tha API to experiment with K8s conversion webhooks.


The tutorial page:
https://book.kubebuilder.io/cronjob-tutorial/cronjob-tutorial
Source code of the tutorial:
https://github.com/kubernetes-sigs/kubebuilder/tree/master/docs/book/src/multiversion-tutorial/testdata/project

### Prerequisites
- go version v1.20.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
Run ./Instructions/create_kind_and_registry.sh which will create a kind cluster and setup docker registry
More info: https://kind.sigs.k8s.io/docs/user/local-registry/


### To deploy controller in local kind cluster

| Deploy changes to K8s cluster            |                                                                                                           |
|------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| Install cert-manager in the kind cluster | kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.2/cert-manager.yaml |
| Build the docker image                   | make docker-build docker-push IMG=localhost:5001/kb:5.0                                                   |
| Deploy the controller,CRD webhooks       | make deploy  IMG=localhost:5001/kb:5.0                                                                    |
| UnDeploy the controller from the cluster | make undeploy                                                                                             |
|                                          |                                                                                                           |


### Testing
Apply v2 of the API
kubectl apply -f config/samples/kb_v2_cronjob.yaml
or alternatively v2 of the API
kubectl apply -f config/samples/kb_v1_cronjob.yaml

k get pod --all-namespaces -> find controller pod -> k logs -n <controller-namespace> <pod-name>
k get jobs -> verify jobs are scheduled


###  To Run locally against local kind cluster
1. kind create cluster -n nasko - create a clsuter
2. make install - installs the CRDs in the cluster under .kube/config
3. make run - run the controller
