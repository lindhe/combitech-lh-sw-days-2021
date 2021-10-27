FROM docker.io/golang:1.17-alpine as builder
WORKDIR /build
COPY main.go go.mod ./
RUN ["go", "build", "-o", "simple-www"]

FROM docker.io/golang:1.17-alpine
WORKDIR /app
COPY --from=builder /build/simple-www ./
COPY index.html.gotmpl ./
EXPOSE 80
ENTRYPOINT ["/app/simple-www"]
