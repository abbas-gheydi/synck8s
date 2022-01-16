# SYNCK8S
## About The Project
synck8s can sync deployments beetween two Kubernetes clusters.   
### How to build
```
go get ./..
go build -o synck8s ./
```
### How to Use
edit config.yml and customize it to meet your needs:   
```yaml
# Kubeconfig address for each cluster. where 'Source' is your main k8s cluster and
#'Destination' is k8s cluster that you want to sync image tags from main cluster.' 
Kubeconfig:
        Source: /root/.kube/master
        Destination: /root/.kube/slave
# 'Deploymetns:' is your deployments that you want to sync image tags.
# 'deployment:' is your kubernetes deployment name.
# 'srcNamespace' is 'deployment' namespace in MAIN kuberntes cluster
# 'dstNamespace' is 'deployment' namespace in TARGET kuberntes cluster
Deployments:
        - backend:
          deployment:   "api"
          srcNamespace: "backend"
          dstNamespace: "backend"

        - frontend:
          deployment:   "website"
          srcNamespace: "frontend"
          dstNamespace: "frontend-slave"
```

### How to Use it:
just run it:   
```
./synck8s
```
output:   
```
Updated deployment api registry.mycompany.com/api:v43
Updated deployment website registry.mycompany.com/website:cd8e1acd

```
