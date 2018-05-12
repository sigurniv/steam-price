FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 8081

ADD steamgame /app/steamgame/steamgame
ADD config.yaml /etc/steamgame/config.yaml

WORKDIR /app/steamgame
#RUN chmod +x /app/nats-chat
CMD ["/app/steamgame/steamgame", "--config", "/etc/steamgame"]