kubectl create configmap experiment-config  --from-file=nginx.conf

kubectl apply -f nginx_volume.yaml
kubectl get pv nginx-pv-volume
kubectl apply -f nginx_vlm_claim.yaml
kubectl get pvc nginx-pv-claim

kubectl create -f dictionary_server.yaml
kubectl create -f echo_server_2.yaml
kubectl create -f experimenet_web_config.yaml
kubectl create -f ambassador_pod.yaml

kubectl create -f ambassador_srvc.yaml
kubectl create -f ingress_nginx_2.yaml

minikube addons enable ingress
minikube tunnel

-  curl http://localhost/WATER VIOLET

kubectl expose pod ambassador --type=LoadBalancer --name=ambassador-service

kubectl describe pod ambassador


## links

- https://kubernetes.io/docs/tutorials/stateless-application/expose-external-ip-address/
- https://kubernetes.io/docs/tasks/configure-pod-container/configure-persistent-volume-storage/
- https://kubernetes.io/docs/tasks/configure-pod-container/assign-memory-resource
- https://minikube.sigs.k8s.io/docs/tutorials/ambassador_ingress_controller/
- https://github.com/twitter/twemproxy
- https://raw.githubusercontent.com/adambom/dictionary/master/dictionary.json
- http://nginx.org/en/docs/http/ngx_http_upstream_module.html#upstream




https://gist.github.com/chukaofili/d0a6713734d0953ce1ce667958464edb

