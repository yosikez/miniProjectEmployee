FROM postgres:15.2
WORKDIR /docker-entrypoint-initdb.d/
COPY ./init-user-db.sh .


FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .


FROM alpine:3
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
EXPOSE 8000
CMD ["./main"]