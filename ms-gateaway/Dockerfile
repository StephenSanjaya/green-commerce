FROM golang:1.21-alpine as builder

RUN apk add --no-cache git

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

EXPOSE 8081

CMD [ "./main" ]