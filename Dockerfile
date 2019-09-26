######################
#       Working      #
######################
FROM golang:alpine
EXPOSE 3000
RUN apk update && apk add --no-cache git

WORKDIR /go/src/github.com/sielerjunjor/framework-api
RUN go get github.com/gorilla/mux go.mongodb.org/mongo-driver ; exit 0
COPY . .
RUN go install
CMD ["framework-api"]
# docker run --rm --net="host" --name golangba golangba

##########
# Step1  #
##########
#FROM golang:latest AS builder
#WORKDIR /go/src/github.com/sielerjunjor/framework-api
#RUN go get github.com/gorilla/mux go.mongodb.org/mongo-driver ; exit 0
#COPY . .
#RUN go build -o /go/src/github.com/sielerjunjor/framework-api


##########
# Step2  #
##########
#FROM scratch
#EXPOSE 3000
#COPY --from=builder /go/src/github.com/sielerjunjor/framework-api /go/src/github.com/sielerjunjor/framework-api
# Run the Frameworks binary.
#ENTRYPOINT ["/go/src/github.com/sielerjunjor/framework-api"]
