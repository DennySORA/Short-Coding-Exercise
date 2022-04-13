FROM golang:1.17-alpine3.14 as builder

ENV GO111MODULE=on
WORKDIR /app
RUN apk add build-base
COPY ./short /app
RUN go build -o short .

FROM alpine:3.14

WORKDIR /app
COPY --from=builder /app/short /app/app.env /app/short.sql /app/ 
CMD ["./short"]
