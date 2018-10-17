FROM golang

RUN mkdir -p /go/src/github.com/jscastelblancoh/statistic_service

ADD . /go/src/github.com/jscastelblancoh/statistic_service

RUN curl https://glide.sh/get | sh
#RUN go get  github.com/canthefason/go-watcher
#RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN cd /go/src/github.com/jscastelblancoh/statistic_service && glide install

#ENTRYPOINT  watcher -run github.com/jscastelblancoh/statistic_service/statistic/cmd -watch github.com/jscastelblancoh/statistic_service/statistic
ENTRYPOINT go run src/github.com/jscastelblancoh/statistic_service/statistic/cmd/main.go