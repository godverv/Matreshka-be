FROM golang:latest as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 go build -o /deploy/server/matreshka-be ./cmd/matreshka-be/main.go

FROM alpine

WORKDIR /app
COPY --from=builder ./deploy/server/ .
COPY --from=builder /app/config/config.yaml ./config/config.yaml

EXPOSE 80

ENTRYPOINT ["./matreshka-be"]