FROM golang:1.21 as builder
WORKDIR /code

COPY go.mod go.sum ./
COPY ./cmd/api/main.go ./cmd/api/main.go
COPY ./pkg ./pkg

# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -gcflags="${SKAFFOLD_GO_GCFLAGS}" -trimpath -o api ./cmd/api/main.go
# RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -trimpath -o meta ./cmd/meta/main.go

FROM gcr.io/distroless/base:latest
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single
COPY --from=builder /code/api .
ENTRYPOINT ["./api"]