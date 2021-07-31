helm repo add harbor https://helm.goharbor.io
helm install harbor harbor/harbor --version 1.5.2 -f values.yaml -n harbor