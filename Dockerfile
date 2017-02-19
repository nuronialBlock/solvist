FROM golang:1.7

# install these before ADD or COPY command to take advantage of caching
RUN apt-get update && apt-get install -y netcat

RUN mkdir -p /go/src/github.com/nuronialBlock/solvist/solvist/
WORKDIR /go/src/github.com/nuronialBlock/solvist/solvist
ADD . /go/src/github.com/nuronialBlock/solvist/solvist
RUN go get ./...
EXPOSE 8080
ENV MONGO_URL mongodb://mongo:27017/solvist
#ENTRYPOINT ["/bin/bash"]