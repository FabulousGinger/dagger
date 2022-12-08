FROM --platform=linux/amd64 golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /dagger-api

EXPOSE 8080

CMD [ "/dagger-api" ]