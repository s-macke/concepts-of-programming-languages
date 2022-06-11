FROM golang:1.18.2

WORKDIR /go/src/app
COPY . .

RUN cd mail && go build
RUN cd plugin && go build -buildmode=plugin plugin.go
RUN cd client && go build client.go

CMD ["/go/src/app/client/client"]
