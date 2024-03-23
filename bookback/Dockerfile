FROM golang:1.22-alpine as builder
WORKDIR /app

COPY ../api_gateway ./api_gateway
COPY ../auth ./auth
COPY ../bookback ./bookback
COPY ../go_pkg ./go_pkg
COPY ../go.work ./go.work

RUN go build -o /app/main ./bookback/cmd/bookback

FROM alpine:3.12
WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /app/bookback/config /app/config

CMD [ "/app/main" ]
