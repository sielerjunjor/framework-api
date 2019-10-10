######################
#       Working      #
######################
FROM scratch
EXPOSE 3000
#RUN apk update && apk add --no-cache git

#WORKDIR /go/src/github.com/sielerjunjor/framework-api
WORKDIR /usr/bin
COPY config.json config.json
COPY app app
#RUN go get github.com/gorilla/mux go.mongodb.org/mongo-driver ; exit 0
#COPY . .
#RUN go install
CMD ["app"]
# docker run --rm --net="host" --name golangba golangba

##########
# Step1  #
##########
#FROM golang:latest AS builder
#WORKDIR /go/src/github.com/sielerjunjor/framework-api
#RUN go get github.com/gorilla/mux go.mongodb.org/mongo-driver ; exit 0
#COPY . .
#RUN go install
#RUN go build -o /go/bin/hello


##########
# Step2  #
##########
#FROM scratch AS Production
#EXPOSE 3000
# Copy our static executable.
#COPY --from=builder /go/bin/hello /go/bin/hello

# Run the hello binary.
#CMD ["./go/bin/hello"]
