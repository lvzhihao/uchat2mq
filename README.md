# uchat2mq
uchat api receive to mq
```
go get github.com/lvzhihao/uchat2mq
```

# docker
```
docker pull edwinlll/uchat2mq
```

## deploy
```
mkdir ~/project/uchat2mq
curl https://raw.githubusercontent.com/lvzhihao/uchat2mq/master/docker-compose.yml -o docker-compose.yml
curl https://raw.githubusercontent.com/lvzhihao/uchat2mq/master/.uchat2mq.yaml.sample -o .uchat2mq.yaml
```

## config
Modify the configuration in .uchat2mq.yaml

## migrate
```
sudo docker-compose run --rm uchat2mq go-wrapper run migrate
```

## run
```
sudo docker-compose up -d
```

## logs
```
sudo docker-compose logs -ft
```

## env
 * TZ default: Asia/Shanghai
 * DEBUG default false
