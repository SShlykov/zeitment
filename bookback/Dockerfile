FROM golang:1.21-rc-alpine as server
WORKDIR /app

COPY ../go.mod .
RUN go mod download

COPY .. .

RUN go build -o /app/main ./cmd/bookback

EXPOSE $PORT

CMD [ "/app/main" ]
