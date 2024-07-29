FROM golang:1.22 AS builder

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client netcat-openbsd

RUN chmod +x wait-for-postgres.sh
RUN chmod +x wait-for-it.sh

RUN go mod download
RUN go build -o talkBoard ./cmd/main.go

CMD ["./talkBoard"]
