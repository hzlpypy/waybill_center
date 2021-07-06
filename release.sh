#!/bin/sh

sudo cp -r ~/.ssh ./docker
sudo cp -r ~/.gitconfig ./docker

LOCAL_IMAGES="waybill_center"
LOCAL_VERSION="1.0.1"

REMOTE_SERVER="waybill_center"
REMOTE_VERSION="1.0.1"
REMOTE_LOGINNAME="1761784585@qq.com"
REMOTE_PATH="registry.cn-hangzhou.aliyuncs.com/k8s_demo1"

#sudo docker rmi -f ${LOCAL_IMAGES}:${LOCAL_VERSION}
#sudo docker rmi -f ${REMOTE_PATH}${REMOTE_SERVER}:${REMOTE_VERSION}

# docker imagesdock
#sudo docker build -t ${LOCAL_IMAGES}:${LOCAL_VERSION} .
sudo docker tag ${LOCAL_IMAGES}:${LOCAL_VERSION} ${REMOTE_PATH}/${REMOTE_SERVER}:${REMOTE_VERSION}

sudo docker login --username=${REMOTE_LOGINNAME} registry.cn-hangzhou.aliyuncs.com
sudo docker push ${REMOTE_PATH}/${REMOTE_SERVER}:${REMOTE_VERSION}

## localhost run
#docker rm ework-server
#docker run -it -p 5010:5009 --name="waybill_center" waybill_center:latest
#docker run -d --name logrus  -p 9201:9200 -p 9301:9300 -e "discovery.type=single-node" elastic/elasticsearch:6.7.1
