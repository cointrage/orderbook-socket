FROM golang:latest 

EXPOSE 3009

RUN mkdir /app 
ADD . /app/ 

WORKDIR /app 

RUN go get github.com/cointrage/orderbook-socket/orderbook
RUN go get github.com/gogo/protobuf/io

RUN go build -o main . 

CMD ["/app/main"]
