
* https://artifacthub.io/packages/helm/bitnami/kafka

helm repo add bitnami https://charts.bitnami.com/bitnami

helm install kafka-release bitnami/kafka

NAME: kafka-release
LAST DEPLOYED: Wed Dec  8 22:06:18 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
CHART NAME: kafka
CHART VERSION: 14.4.3
APP VERSION: 2.8.1

** Please be patient while the chart is being deployed **

Kafka can be accessed by consumers via port 9092 on the following DNS name from within your cluster:

    kafka-release.default.svc.cluster.local

Each Kafka broker can be accessed by producers via port 9092 on the following DNS name(s) from within your cluster:

    kafka-release-0.kafka-release-headless.default.svc.cluster.local:9092

To create a pod that you can use as a Kafka client run the following commands:

    kubectl run kafka-release-client --restart='Never' --image docker.io/bitnami/kafka:2.8.1-debian-10-r57 --namespace default --command -- sleep infinity
    kubectl exec --tty -i kafka-release-client --namespace default -- bash

    PRODUCER:
        kafka-console-producer.sh \
            --broker-list kafka-release-0.kafka-release-headless.default.svc.cluster.local:9092 \
            --topic test

    CONSUMER:
        kafka-console-consumer.sh \
            --bootstrap-server kafka-release.default.svc.cluster.local:9092 \
            --topic test \
            --from-beginning
