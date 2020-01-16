FROM alpine:3.10

LABEL maintainer="jack <yongjie.zhang@henganpros.com>"

ENV ADDRESS :20020

RUN apk --no-cache add ca-certificates

COPY main /usr/local/bin/

CMD ["main"]

EXPOSE 20020