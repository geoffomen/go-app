FROM golang:1.20.1 AS builder
ARG APP_NAME
ENV APP_NAME=${APP_NAME}
COPY . /opt/${APP_NAME}
RUN cd /opt/${APP_NAME} && make app args_appname=${APP_NAME}


FROM debian:11
ARG APP_NAME
ENV APP_NAME=${APP_NAME}
RUN sed -i "s@http://deb.debian.org/debian@http://mirror.sjtu.edu.cn/debian@g" /etc/apt/sources.list \
    && apt-get update && apt-get install -y ca-certificates 
COPY --from=builder /opt/${APP_NAME}/build/temp/${APP_NAME}.o /opt/${APP_NAME}/
COPY ./assets /opt/${APP_NAME}/assets
COPY ./deployments/container/docker_compose_${APP_NAME}.yml /opt/${APP_NAME}/compose.yml
COPY ./configs /opt/${APP_NAME}/configs
COPY ./deployments /opt/${APP_NAME}/deployments
COPY ./scripts /opt/${APP_NAME}/scripts
COPY ./website /opt/${APP_NAME}/website
COPY ./deployments/container/init.sh /
WORKDIR /opt/${APP_NAME}
CMD ["/init.sh"]