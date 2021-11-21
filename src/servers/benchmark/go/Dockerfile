FROM golang:1.17
WORKDIR /root/
COPY server.go ./
RUN go mod init server
RUN CGO_ENABLED=0 GOOS=linux go build

FROM scratch
WORKDIR /
COPY --from=0 /root/server ./
COPY cert.crt cert.key ./
CMD ["./server"]
