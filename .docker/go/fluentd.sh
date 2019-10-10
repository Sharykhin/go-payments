#!/usr/bin/env bash

for i in {1..10}; do
    status_code=$(curl --write-out %{http_code} --silent --output /dev/null http://fluentd:24220/api/plugins.json)
    if [[ "$status_code" -eq 200 ]] ; then
        echo "Fluentd is up and running."
        exit 0
    else
        echo "Fluentd is not ready yet wait 5 seconds"
        sleep 5
    fi
done
exit 1
