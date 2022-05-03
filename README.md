# provider-gps-dummy

`provider-gps-dummy` is a dummy [Crossplane](https://crossplane.io/) provider providing fake GPS data for a vessel with a given IMO.

## Developing

Install CRDs:
```
# generate CRDs
make generate
# install CRDs
kubectl apply -f package/crds/
# check
kubectl get crds | grep gps
```

Configure the Provider:
```
kubectl apply -f examples/provider/config.yaml
```

Run the Provider Controller locally:
```
make run
```

Install example
```
kubectl create -f examples/sample/vesselgpstype-9259501.yaml
```

Delete Examples
```
# delete
kubectl delete -f examples/sample/vesselgpstype-9259501.yaml

# force delete
kubectl patch vesselgpstypes.sample.gps.crossplane.io vessel-imo-9259501 -p '{"metadata":{"finalizers": []}}' --type=merge
```

## TODOs
* Why do I need to force delete?
* Build and publish provider package
* More dummy data
* delete storeconfig