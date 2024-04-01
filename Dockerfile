FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk --no-cache add bash git make gettext

#dependencies
COPY go.* ./
RUN go mod download

COPY ./ ./

RUN go build -o ./bin/article-service cmd/main.go

FROM alpine AS runner

COPY --from=builder /app/bin/article-service  /
COPY internal/config/config_local.yaml config.yaml

CMD ["/article-service --path=config.yaml"]