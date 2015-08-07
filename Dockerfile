FROM golang
RUN go get gopkg.in/mgo.v2
COPY ./simple-golang-app /go/src/simple-golang-app
WORKDIR /go/src/simple-golang-app
CMD go run main.go
EXPOSE 8080