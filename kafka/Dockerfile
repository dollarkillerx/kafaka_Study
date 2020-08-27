FROM anapsix/alpine-java

LABEL author="Koma<komazhang@foxmail.com>"

ENV KAFKA_VERSION 2.5.0
ENV SCALA_VERSION 2.12
ENV KAFKA_INSTALL_PATH /usr/local

#设置软件源
RUN echo 'https://mirrors.ustc.edu.cn/alpine/v3.6/community/' >  /etc/apk/repositories \
    && echo 'https://mirrors.ustc.edu.cn/alpine/v3.6/main/' >> /etc/apk/repositories

#设置时区
RUN apk update && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Hong_Kong /etc/localtime \
    && echo 'Asia/Hong_Kong' > /etc/timezone \
    && apk del tzdata

RUN apk add --update unzip wget curl supervisor

VOLUME ["/kafka"]

COPY scripts/install.sh /tmp/install.sh
RUN chmod a+x /tmp/install.sh && sync && /tmp/install.sh

ENV PATH ${PATH}:/opt/kafka/bin

COPY conf/supervisord.conf /etc/supervisord.conf
COPY conf/kafka-server.properties /opt/kafka/config/server.properties

COPY scripts/start.sh /usr/bin/start-kafka.sh
RUN chmod a+x /usr/bin/start-kafka.sh

CMD ["start-kafka.sh"]