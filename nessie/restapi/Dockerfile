FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY nessie/restapi/go.mod nessie/restapi/go.sum ./
RUN go mod download

# Cache busting: Rebuild the image when the source files change
ARG UNIQUE=default
RUN echo "Cache bust: $UNIQUE"

COPY nessie/restapi/main.go nessie/restapi/index.html nessie/restapi/nessie-openapi.yaml ./
RUN CGO_ENABLED=0 go build -o nessie-restapi main.go

FROM alpine:3.23 AS runner

WORKDIR /app

COPY --from=builder /app/nessie-restapi .
COPY nessie/restapi/index.html nessie/restapi/nessie-openapi.yaml ./

CMD ["./nessie-restapi"]