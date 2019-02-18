FROM golang:latest as builder

WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 go build -a -o main .


FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/src/app/main .

CMD ["/root/main"]
