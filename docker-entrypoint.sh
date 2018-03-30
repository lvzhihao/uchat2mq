#!/bin/sh

# default timezone
if [ ! -n "$TZ" ]; then
    export TZ="Asia/Shanghai"
fi

# set timezone
ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
echo $TZ > /etc/timezone 

# k8s config  switch
if [ -f "/usr/local/uchat2mq/config/.uchat2mq.yaml" ]; then
    ln -s  /usr/local/uchat2mq/config/.uchat2mq.yaml /usr/local/uchat2mq/.uchat2mq.yaml
fi

# version
/usr/local/uchat2mq/uchat2mq version

# apply config
echo "===config start==="
cat /usr/local/uchat2mq/.uchat2mq.yaml
echo "====config end===="

# run command
if [ ! -z "$1" ]; then
    /usr/local/uchat2mq/uchat2mq $@
fi
