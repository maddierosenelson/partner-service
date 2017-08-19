FROM golang:1.8.1-alpine

EXPOSE 80

COPY partner_service /go/bin/partner-service

CMD partner-service --grpcAddr :8080 --httpAddr :80