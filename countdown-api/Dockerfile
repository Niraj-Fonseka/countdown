FROM golang:alpine as builder

COPY . /app
WORKDIR /app

RUN go mod download 

RUN CGO_ENABLED=0  GOOS=linux go build -a -installsuffix cgo -o app .




FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]