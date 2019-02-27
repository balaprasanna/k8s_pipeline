# Step 1.1:
- Run the following on Main Node
```
kubeadm init --apiserver-advertise-address $(hostname -i)
```

# Step 1.2:
- Initialize cluster networking
```
kubectl apply -n kube-system -f \
    "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 |tr -d '\n')"
```

# Step 2:
- Wait until you see the kubeadm join tokens, something like the one showed below
```
kubeadm join 192.168.0.18:6443 --token nt8ywh.899sniuwhewtaypu --discovery-token-ca-cert-hash sha256:fe31ef21d38fca421f79efe71c9026cba440848abfb64773464590fcd8533417
```

# Step 3:
- Add a new instance and run the kubeadm join command with the token captured above
```
kubeadm join 192.168.0.18:6443 --token nt8ywh.899sniuwhewtaypu --discovery-token-ca-cert-hash sha256:fe31ef21d38fca421f79efe71c9026cba440848abfb64773464590fcd8533417
```

# Step 4:
- Go to Master Node and check if the new nodes are added to the cluster
```
$ kubectl get nodes
NAME      STATUS    ROLES     AGE       VERSION
node1     Ready     master    47m       v1.11.3
node2     Ready     <none>    40m       v1.11.3
```


