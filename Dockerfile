FROM golang:1.21

RUN apt-get update && apt-get install -y gcc sqlite3 libsqlite3-dev

WORKDIR /app
COPY . .

ENV CGO_ENABLED=1
RUN go build -o app .

CMD ["./app"]
