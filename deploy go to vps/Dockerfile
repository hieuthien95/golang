FROM golang:1.11.4-alpine
WORKDIR /go/src/go-module
COPY . /go/src/go-module/
RUN go build -o main . 
CMD ["./main"]

