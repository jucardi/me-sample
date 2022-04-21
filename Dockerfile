FROM alpine:latest

RUN apk update && \
    apk upgrade && \
    apk add ca-certificates curl --update --no-cache

EXPOSE 8080


COPY fixtures/* /fixtures/
COPY bin/{{.service_name}}-linux /{{.service_name}}

CMD ["/{{.service_name}}"]

