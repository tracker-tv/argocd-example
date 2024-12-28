FROM golang:1.23.4 AS builder

WORKDIR /app
COPY go.* ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/build/argocd-example ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /tmp/build/argocd-example /

EXPOSE 8080

CMD ["/argocd-example"]
