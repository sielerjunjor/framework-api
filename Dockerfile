FROM golang:latest

WORKDIR /go/src/github.com/sielerjunjor/framework-api
COPY . .

RUN go get -v
RUN go install

EXPOSE 3000

CMD ["framework-api"]