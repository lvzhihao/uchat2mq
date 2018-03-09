#!/bin/sh

# default timezone
if [ ! -n "$TZ" ]; then
    export TZ="Asia/Shanghai"
fi

# set timezone
ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
echo $TZ > /etc/timezone 

# k8s config  switch
if [ -f "/usr/local/uchat2mq/config/uchat2mq.yaml.base64" ]; then
    base64 -d /usr/local/uchat2mq/config/uchat2mq.yaml.base64 > /usr/local/uchat2mq/.uchat2mq.yaml
fi

# apply config
echo "===start==="
cat /usr/local/uchat2mq/.uchat2mq.yaml
echo "====end===="

# run command
if [ ! -z "$1" ]; then
    /usr/local/uchat2mq/uchat2mq $@
fi
