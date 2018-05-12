FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 8080

ADD currency_rate /app/currency_rate/currency_rate
ADD config.yaml /etc/currency_rate/config.yaml

WORKDIR /app/currency_rate
#RUN chmod +x /app/nats-chat
CMD ["/app/currency_rate/currency_rate", "--config", "/etc/currency_rate"]