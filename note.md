元数据存储格式

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: clusterId
  labels:
    metadata: true
   annotations:
    displayName: "集群名"
    city: "城市"
    district: "区"
type: generic
data:
  kubeconfig: <base64-encoded-kubeconfig>
```