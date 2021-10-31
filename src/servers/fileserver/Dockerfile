FROM golang:1.17
#WORKDIR /go/src/github.com/alexellis/href-counter/
#RUN go get -d -v golang.org/x/net/html
WORKDIR /root/
COPY main.go ./
RUN go mod init fileserver
RUN CGO_ENABLED=0 GOOS=linux go build

FROM scratch
WORKDIR /
COPY --from=0 /root/fileserver ./
CMD ["./fileserver"]
 m