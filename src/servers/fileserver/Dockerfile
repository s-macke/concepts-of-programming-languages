FROM golang:1.19
WORKDIR /root/
COPY main.go ./
RUN go mod init fileserver
RUN CGO_ENABLED=0 GOOS=linux go build

FROM scratch
WORKDIR /
COPY --from=0 /root/fileserver ./
CMD ["./fileserver"]
