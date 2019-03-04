FROM golang:latest

RUN mkdir /go/src/api
RUN go get -u github.com/gin-gonic/gin

COPY main.go /go/src/api

CMD ["go", "run", "/go/src/api/main.go"]