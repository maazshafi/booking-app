FROM golang:1.16-alpine

WORKDIR /app

ENV GO111MODULE=on
COPY go.mod .
RUN go mod download

COPY . .


RUN go build -o /booking-app

CMD ["/booking-app"]