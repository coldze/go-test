FROM golang:1.10.1

RUN go get github.com/coldze/go-test
WORKDIR /go/src/github.com/coldze/go-test
RUN go install ./...

CMD ["go-test"]