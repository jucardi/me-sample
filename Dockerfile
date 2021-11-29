FROM alpine:latest

RUN apk update && \
    apk upgrade && \
    apk add ca-certificates curl --update --no-cache

EXPOSE 8080


RUN mkdir /root/.ssh

COPY fixtures/known_hosts /root/.ssh/known_hosts
COPY fixtures/default-docker.yml /fixtures/config.yml
COPY bin/ms-sample-linux /ms-sample

CMD ["/ms-sample"]

