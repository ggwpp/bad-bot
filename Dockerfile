FROM golang:1.11.5-alpine as builder
LABEL maintainer="Grean-Developers-Family"

RUN apk add --no-cache git \
    && go get github.com/line/line-bot-sdk-go/linebot
WORKDIR /go/src/app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /go/src/app
COPY --from=builder /go/src/app/app .
CMD [ "./app" ]
EXPOSE 6969