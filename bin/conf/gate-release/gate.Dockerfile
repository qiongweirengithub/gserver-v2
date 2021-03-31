# Docker image for springboot file run
# VERSION 0.0.1
# Author: eangulee
# 基础镜像使用java
FROM alpine:3.5
# 作者
MAINTAINER dika <dika@gmail.com>
# VOLUME 指定了临时文件目录为/tmp。
# 其效果是在主机 /var/lib/docker 目录下创建了一个临时文件，并链接到容器的/tmp
VOLUME /tmp 

WORKDIR ～

# 将jar包添加到容器中并更名为app.jar
ADD gserverv2 /usr/bin/gserverv2 
ADD server.json server.json
ADD env.json env.json

ENTRYPOINT ["gserverv2","-conf","server.json","-pid", "gateconnectionsvc"]
# ENTRYPOINT ["gserverv2","-consul_addr", "172.17.0.1:8500", "-nats_addr", "nats://172.17.0.1:4222", "-conf","server.json"]

