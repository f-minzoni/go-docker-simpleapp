FROM golang
RUN go get gopkg.in/mgo.v2
COPY simple-golang-app /staticbinary
ENTRYPOINT ["/staticbinary"]
EXPOSE 8080