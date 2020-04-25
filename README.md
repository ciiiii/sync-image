# sync-image

1. only support docker.io、gcr.io、quay.io.
2. k8s.gcr.io should convert to gcr.io/google-containers format

```json
{
  "docker.io": {
    "library": {
      "nginx": ["1.17"]
    }
  },
  "quay.io": {
      "coreos": {
        "etcd": ["v3.3.12"]
    }
  },
  "gcr.io":{
    "google-containers": {
      "kube-proxy": ["v1.17.4"]
    }
  }
}
```