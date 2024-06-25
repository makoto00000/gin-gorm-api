FROM golang:1.22.4-alpine

WORKDIR /app

COPY app/go.mod .
COPY app/go.sum .

RUN go mod download
RUN apk update
RUN apk add --no-cache gcc
RUN apk add --no-cache musl-dev

COPY ./app .

CMD ["go", "run", "./main.go"]
