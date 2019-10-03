FROM golang:latest
WORKDIR /app
COPY *.go /app
RUN go build -o server *.go

FROM ubuntu:latest

WORKDIR /app
EXPOSE 8080
USER root

COPY --from=0 /app/server /app/server

CMD ["/app/server"]

