FROM golang:1.16-buster
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o webApp .
WORKDIR /dist
RUN cp /build/webApp .
EXPOSE 3000
ENTRYPOINT ["/dist/webApp"]
