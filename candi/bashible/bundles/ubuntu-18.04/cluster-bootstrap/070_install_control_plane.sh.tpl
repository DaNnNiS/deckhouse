kubeadm init phase certs all --config /var/lib/bashible/kubeadm/config.yaml
kubeadm init phase kubeconfig all --config /var/lib/bashible/kubeadm/config.yaml
kubeadm init phase etcd local --config /var/lib/bashible/kubeadm/config.yaml -k /var/lib/bashible/kubeadm/kustomize
kubeadm init phase control-plane all --config /var/lib/bashible/kubeadm/config.yaml -k /var/lib/bashible/kubeadm/kustomize
# TODO возможно здесь понадобиться дождаться, пока control-plane поднимется
kubeadm init phase mark-control-plane --config /var/lib/bashible/kubeadm/config.yaml

# Upload pki for deckhouse
kubectl -n kube-system create secret generic d8-pki \
  --from-file=ca.crt=/etc/kubernetes/pki/ca.crt \
  --from-file=ca.key=/etc/kubernetes/pki/ca.key \
  --from-file=sa.pub=/etc/kubernetes/pki/sa.pub \
  --from-file=sa.key=/etc/kubernetes/pki/sa.key \
  --from-file=front-proxy-ca.crt=/etc/kubernetes/pki/front-proxy-ca.crt \
  --from-file=front-proxy-ca.key=/etc/kubernetes/pki/front-proxy-ca.key \
  --from-file=etcd-ca.crt=/etc/kubernetes/pki/etcd/ca.crt \
  --from-file=etcd-ca.key=/etc/kubernetes/pki/etcd/ca.key

# Setup kubectl for root user
if [ ! -f /root/.kube/config ]; then
  mkdir -p /root/.kube
  ln -s /etc/kubernetes/admin.conf /root/.kube/config
fi
