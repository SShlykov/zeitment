FROM golang:1.22-alpine as server
WORKDIR /app

COPY ../go.mod .
RUN go mod download

COPY .. .

RUN go build -o /app/main ./api_gateway/cmd/api_gateway

CMD [ "/app/main" ]
