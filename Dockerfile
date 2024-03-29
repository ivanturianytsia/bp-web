FROM alpine

RUN apk --update upgrade && \
    apk --update add ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

RUN mkdir /web

WORKDIR /web

ADD ./bin/web_alpine /web/web
ADD ./templates /web/templates

CMD ["/web/web","-addr=:8080"]
