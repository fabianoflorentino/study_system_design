FROM golang:alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN go build -o /go/bin/system_design cmd/system_design/main.go

FROM alpine:3.20.3 as system_design

COPY --from=builder /go/bin/system_design /usr/local/bin/system_design

RUN apk update && apk upgrade --no-cache

ENTRYPOINT [ "/bin/sh" ]
CMD ["-c", "while true; do echo 'Im running...'; sleep 300; done"]
