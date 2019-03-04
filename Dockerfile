FROM golang:latest

RUN mkdir /go/src/reverpro

COPY main.go /go/src/reverpro

CMD ["go", "run", "/go/src/reverpro/main.go"]