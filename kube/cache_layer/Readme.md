
kubectl create -f dictionary-deploy.yaml

kubectl create -f dictionary-service.yaml

kubectl create configmap varnish-config --from-file=default.vcl

kubectl create -f varnish-deploy.yaml

kubectl create -f varnish-service.yaml

kubectl create -f cache-ingress.yaml


https://raw.githubusercontent.com/adambom/dictionary/master/dictionary.json

curl http://localhost/DEMOCRATIST

curl -w "@curl-format.txt" -o /dev/null -s http://localhost/CARBURIZATION


 curl -w "@curl-format.txt" -o /dev/null -s http://localhost/DEMOCRATIST

 curl_time http://localhost/DEMOCRATIST
