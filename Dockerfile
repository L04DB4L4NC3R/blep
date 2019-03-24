FROM golang

RUN mkdir -p /go/src/github.com/angadsharma1016/mormon

ADD . /go/src/github.com/angadsharma1016/mormon
WORKDIR /go/src/github.com/angadsharma1016/mormon/server

#RUN go get -t -v ./...
RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher
EXPOSE 3000
ENTRYPOINT watcher -run github.com/angadsharma1016/mormon/server/ -watch github.com/angadsharma1016/mormon/server