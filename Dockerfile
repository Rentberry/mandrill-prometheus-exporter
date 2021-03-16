FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY ./mandrill-prometheus-exporter .
ENTRYPOINT ["./mandrill-prometheus-exporter"]

EXPOSE 9861
