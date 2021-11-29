kubectl create configmap experiment-config  --from-file=nginx.conf

kubectl apply -f nginx_volume.yaml
kubectl get pv nginx-pv-volume
kubectl apply -f nginx_vlm_claim.yaml
kubectl get pvc nginx-pv-claim

kubectl create -f dictionary_server.yaml
kubectl create -f echo_server_2.yaml
kubectl create -f experimenet_web_config.yaml
kubectl create -f ambassador_pod.yaml

kubectl expose pod ambassador --type=LoadBalancer --name=ambassador-service


https://gist.github.com/chukaofili/d0a6713734d0953ce1ce667958464edb

