FROM golang:alpine

LABEL maintainer="manuel.fuchs49@gmail.com"

WORKDIR /go/src

COPY .env /go/src
COPY /cmd /go/src/backend

RUN go mod init github.com/manuelfuchs/go-products.git
RUN go get -v github.com/gorilla/mux
RUN go get -v github.com/lib/pq
RUN go get -v github.com/joho/godotenv

RUN go build -o /usr/backend ./...

EXPOSE 80

CMD ["/usr/backend"]