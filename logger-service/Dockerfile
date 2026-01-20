FROM golang:1.24.12-alpine3.23 as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o loggerServiceApp ./cmd/api
RUN chmod +x /app/loggerServiceApp

FROM alpine:latest as runner
RUN mkdir /app
COPY --from=builder /app/loggerServiceApp /app
CMD ["/app/loggerServiceApp"]

