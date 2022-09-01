# syntax=docker/dockerfile:1

FROM golang:1.18-bullseye
WORKDIR /app

COPY go.* ./
COPY *.go ./
COPY utils/ ./
COPY tmpl/ ./

RUN go mod download
RUN go build .

EXPOSE 8080

CMD [ "/sn" ]