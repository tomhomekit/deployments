### 在k8s中搭建 mongo shared

搭建 `mongo shared` 集群, 目标:

```shell
# shared1
$ k exec -it shared1-0 bash
root@shared1-0:/# mongo --port 27018
> 
kubectl -n mongodb \
  create secret generic ops-manager-admin-key \
  --from-file="user=/Users/zhangmengjin/.ssh/id_rsa.pub" \
  --from-file="publicApiKey=/Users/zhangmengjin/.ssh/id_rsa"
  
kubectl create configmap mongoshared \
--from-literal="baseUrl=http://mongo.mrj.com" \
--from-literal="projectName=mongoshared" \
--from-literal="orgId=mongoshared" 
```

```yaml
apiVersion: mongodb.com/v1
kind: MongoDB
metadata:
  name: demo-mongodb-cluster-1
  namespace: mongodb
spec:
  members: 3
  version: 4.4.5-ent
  type: ReplicaSet
  authentication:
    enabled: true
    modes: [ "SHA" ]
  opsManager:
    configMapRef:
      name: myconfigmap
  credentials: ops-manager-admin-key
  persistent: true
  podSpec:
    podTemplate:
      spec:
        containers:
          - name: mongodb-enterprise-database
            resources:
              limits:
                cpu: 2
                memory: 1.5G
              requests:
                cpu: 1
                memory: 1G
              persistence:
                single:
                  storage: 10Gi
```