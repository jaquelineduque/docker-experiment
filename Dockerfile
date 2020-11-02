FROM golang:1.15
ENV APP_PORT=8080

WORKDIR /app/server
COPY . /app/server

RUN go build -o /app/bin
CMD ["/app/bin"]