FROM alpine:latest

RUN apk update && apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Yekaterinburg /etc/localtime && \
    echo "Asia/Yekaterinburg" > /etc/timezone && \
    apk del tzdata

COPY notes /
EXPOSE 1323
CMD ["/notes"]