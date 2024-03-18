FROM golang:1.22-alpine as builder
WORKDIR /app

COPY ./api_gateway ./api_gateway
COPY ./auth ./auth
COPY ./bookback ./bookback
COPY ./logger ./logger
COPY ./metrics ./metrics
COPY ./postgres ./postgres
COPY ./scripts ./scripts
COPY ./go.work ./go.work

RUN go build -o /app/main ./api_gateway/cmd/api_gateway

FROM alpine:3.12
WORKDIR /app

COPY --from=builder /app/main /app/main

CMD [ "/app/main" ]
