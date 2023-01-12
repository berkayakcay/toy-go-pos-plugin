TOY POS PLUGIN
---


**Colima Start**

https://github.com/abiosoft/colima
```
colima stop
colima start --cpu 4 --memory 8
```

**K8S Dashboard**

https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
kubectl apply -f zarf/k8s/dashboard/k8s-dashboard-service-account.yaml
kubectl apply -f zarf/k8s/dashboard/k8s-dashboard-cluster-role-binding.yaml
kubectl -n kubernetes-dashboard create token admin-user
```


**Kind Up&Running**

```
make kind-down
make kind-up
make kind-load
(*) make kind-status-sales 
make kind-apply
(*) make kind-logs-sales 
```