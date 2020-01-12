#!/bin/sh

redis_address=$(env | grep REDIS_ADDRESS | cut -f 2 -d =)
host=$(echo ${redis_address} | cut -f 1 -d :)
port=$(echo ${redis_address} | cut -f 2 -d :)

ping() {
    echo quit | nc ${host} ${port} | grep "+OK"
}

check() {
    for i in $(seq 1 10); do
        resp=$(ping)

        if [ ! -z "${resp}" ]; then
            return 0
        fi

        printf "entrypoint: Failed to open connection to redis at %s:%s, retrying in 1s...\n" ${host} ${port}
        sleep 1
    done

    printf "entrypoint: Cannot reach redis at %s:%s, terminating\n" ${host} ${port}
    exit 1
}

printf "entrypoint: Starting listener for redis at %s:%s\n" ${host} ${port}
check
printf "entrypoint: Successfully pinged redis at %s:%s\n" ${host} ${port}

exec $@