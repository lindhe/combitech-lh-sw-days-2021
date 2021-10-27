FROM docker.io/golang:1.17-alpine as builder
WORKDIR /build
COPY main.go go.mod ./
RUN CGO_ENABLED=0 go build -o simple-www -a -ldflags '-extldflags "-static"' .

FROM scratch
WORKDIR /app
COPY --from=builder /build/simple-www ./
COPY index.html.gotmpl ./
EXPOSE 80
ENTRYPOINT ["/app/simple-www"]
