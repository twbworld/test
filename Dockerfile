FROM ubuntu:20.04

LABEL maintainer="twb<1174865138@qq.com><github.com/twbworld>"
LABEL description="构建v2ray-trojan镜像"

WORKDIR /root

# ARG INSTALL=https://raw.githubusercontent.com/twbworld/docker-v2ray-trojan/master/install.sh

ADD install.sh .

RUN set -xe \
        && cd /root \
        # && wget -N --no-check-certificate -q -O install.sh "${INSTALL}" \
        && chmod +x *.sh \
