#!/usr/bin/env bash

for i in {1..10}; do
    status_code=$(curl --write-out %{http_code} -u guest:guest --silent --output /dev/null http://rabbitmq:15672/api/overview)
    if [[ "$status_code" -eq 200 ]] ; then
        echo "RabbitMQ is up and running"
        exit 0
    else
        echo "RabbitMQ is not ready yet wait 5 seconds"
        sleep 5
    fi
done
exit 1