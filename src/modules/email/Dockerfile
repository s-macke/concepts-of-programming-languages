FROM golang:1.19

WORKDIR /go/src/app
COPY . .

RUN cd mail && go build
RUN cd client && go build client.go

CMD ["/go/src/app/client/client"]
