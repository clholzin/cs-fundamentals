kubectl create -f redis-shards-v1.yaml

kubectl create -f redis-service.yaml

kubectl create configmap twem-config --from-file=./nutcracker.yaml

kubectl create -f redis-ambassador-example.yaml

exec into abassador

```
curl -X POST -d "get key\r\n" localhost:6222
```

```
{"service":"nutcracker", "source":"ambassador-example", "version":"0.3.0", "uptime":716, "timestamp":1638084420, "total_connections":1, "curr_connections":1, "redis": {"client_eof":0, "client_err":0, "client_connections":0, "server_ejects":0, "forward_error":0, "fragments":0, "sharded-redis-0.redis:6379": {"server_eof":0, "server_err":0, "server_timedout":0, "server_connections":0, "server_ejected_at":0, "requests":0, "request_bytes":0, "responses":0, "response_bytes":0, "in_queue":0, "in_queue_bytes":0, "out_queue":0, "out_queue_bytes":0},"sharded-redis-1.redis:6379": {"server_eof":0, "server_err":0, "server_timedout":0, "server_connections":0, "server_ejected_at":0, "requests":0, "request_bytes":0, "responses":0, "response_bytes":0, "in_queue":0, "in_queue_bytes":0, "out_queue":0, "out_queue_bytes":0},"sharded-redis-2.redis:6379": {"server_eof":0, "server_err":0, "server_timedout":0, "server_connections":0, "server_ejected_at":0, "requests":0, "request_bytes":0, "responses":0, "response_bytes":0, "in_queue":0, "in_queue_bytes":0, "out_queue":0, "out_queue_bytes":0}}}

```
