FROM golang
ENTRYPOINT ["/staticbinary"]
RUN go get gopkg.in/mgo.v2
COPY ./simple-golang-app /go/src/simple-golang-app
WORKDIR /go/src/simple-golang-app
COPY go-docker-simpleapp /staticbinary
EXPOSE 8080