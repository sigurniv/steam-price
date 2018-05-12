FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 8083

ADD steamprice /app/steamprice/steamprice
ADD config.yaml /etc/steamprice/config.yaml

WORKDIR /app/steamprice
#RUN chmod +x /app/nats-chat
CMD ["/app/steamprice/steamprice", "--config", "/etc/steamprice"]