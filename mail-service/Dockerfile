FROM golang:1.24.12-alpine3.23 as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o mailerApp ./cmd/api
RUN chmod +x /app/mailerApp

FROM alpine:latest as runner
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/mailerApp /app
COPY --from=builder /app/templates /app/templates
CMD ["/app/mailerApp"]

