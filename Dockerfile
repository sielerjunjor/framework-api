FROM golang:latest

EXPOSE 3000

WORKDIR /go/src/github.com/sielerjunjor/framework-api
RUN go get github.com/gorilla/mux go.mongodb.org/mongo-driver ; exit 0

COPY . .

RUN go install

CMD ["framework-api"]

# docker run --rm --net="host" --name golangba golangba