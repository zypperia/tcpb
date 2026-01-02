FROM golang:latest as build

ENV LISTEN_PORT 8081
ENV CONNECT_PORT 8080

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -ldflags="-s -w" -v

FROM gcr.io/distroless/static-debian13
COPY --from=build /app /
ENTRYPOINT ["/tcpb"]
