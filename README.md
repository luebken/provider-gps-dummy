# provider-gps-dummy

`provider-gps-dummy` is a dummy [Crossplane](https://crossplane.io/) provider providing fake GPS data for a vessel with a given IMO.

## Usage

```
# install
kubectl crossplane install provider luebken/provider-gps-dummy:v0.1.4

# configure 
kubectl apply -f examples/provider/config.yaml

# example
kubectl create -f examples/sample/vesselgpstype-9259501.yaml
```

## Developing

```
# generate CRDs
make generate
# install CRDs
kubectl apply -f package/crds/
# check
kubectl get crds | grep gps
# Run the Provider Controller locally:
make run
```

Delete Examples
```
# delete
kubectl delete -f examples/sample/vesselgpstype-9259501.yaml
# force delete
kubectl patch vesselgpstypes.sample.gps.crossplane.io vessel-imo-9259501 -p '{"metadata":{"finalizers": []}}' --type=merge
```

## Relase
* Set env.DOCKER_USR & env.DOCKER_PSW
* Set env.AWS to dummy

* Run Github Action Tag
* Run CI from main branch

## Notes
* Fix: Permission: https://github.com/luebken/provider-gps-dummy/commit/0ef7a396e4d710df6c8a205a349cfad88f3ec25e
* Fix: Test Integration https://github.com/luebken/provider-gps-dummy/commit/83d26617c14c96f7595f33669be3b0c7dd5fb530

## TODOs
* Why do I need to force delete MRs?
* Add more dummy data
* Delete storeconfig