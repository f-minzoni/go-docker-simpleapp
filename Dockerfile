# requires statically linked go binary to be compiled
# to ./go-docker-simpleapp before docker build
FROM scratch
COPY go-docker-simpleapp /staticbinary
ENTRYPOINT ["/staticbinary"]
EXPOSE 8080