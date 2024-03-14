FROM golang:1.22-alpine as server
WORKDIR /app

COPY ../go.mod .
RUN go mod download

COPY .. .

RUN go build -o /app/main ./cmd/bookback

#FROM alpine:latest
#WORKDIR /app
#COPY --from=builder /app/main /app/main
#
#EXPOSE $PORT

CMD [ "/app/main" ]
