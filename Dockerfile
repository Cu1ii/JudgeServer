FROM gleaming/golang1.9.3:env
MAINTAINER cu1

ADD . ./Judge-server
WORKDIR ./Judge-server

ENV GO111MODULE on
ENV CGO_ENABLED 0

RUN mkdir build && cd build && cmake ../Judger && make && make install && cd ../xoj_judgehost

# RUN go mod init && go mod tidy
CMD go build xoj_judgehost