# k8s-secret-generator

Generate random data and save it as a Secret within your Kubernetes cluster.

Running `kubectl apply -f job.yaml` would create a Secret similar to the following:

```
apiVersion: v1
kind: Secret
metadata:
  creationTimestamp: 2018-09-10T17:10:44Z
  name: my-secret
  namespace: default
  resourceVersion: "692264"
  selfLink: /api/v1/namespaces/default/secrets/my-secret
  uid: 741ca711-b51c-11e8-a2b3-42010a8a027d
type: Opaque
data:
  my-key: M1UyaVIreTI1QWpPcGJDTlkwS0kxdHBuTUZXUThCUzlTaWE1VTBaL3hmOD0=
```

## Params
| Flag | Usage |
| ---- | ----- |
| name | Required. The metadata.name for the Secret object |
| key | The key within the Secret data where the generated secret will be saved. Defaults to `data` |
| length | Byte length of the secret. Defaults to 32 |
| base64encode | Encode the random bytes as a base64 string before saving. This is separate to the base64 encoding applied by Kubernetes to store the secret. Default is `false`. |
| namespace | The namespace the Secret will be created in. Defaults to `default`. |

## Authorization
Your default `kubectl` user may not have permission to create the Role and RoleBinding required for the ServiceAccount to create a Secret resource. To apply the yaml as an admin user on GKE run the following commands:

```
gcloud container clusters update my-cluster --enable-basic-auth
PASS=$(gcloud container clusters describe my-cluster | grep password | awk '{ print $2 }')
kubectl config set-credentials admin --username=admin --password=$PASS
kubectl --user=admin apply -f job.yaml
```
