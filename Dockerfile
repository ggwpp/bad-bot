## build
FROM golang:1.11.5-alpine as builder
LABEL maintainer="Grean-Developers-Family"

WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git \
    && go get -d -v
RUN GO_ENABLED=0 GOOS=linux go install .

## app
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /go/bin/app .
CMD [ "/app/app" ]
EXPOSE 6969