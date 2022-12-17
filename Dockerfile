FROM gleaming/golang1.9.3:env
MAINTAINER cu1

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
ENV CGO_ENABLED 0

WORKDIR $GOPATH/src/xoj_judgehost

ADD . $GOPATH/src/xoj_judgehost

RUN mkdir build && cd build && cmake ../JudgerCore && make && make install && cd .. && go mod tidy

RUN cd $GOPATH/src/xoj_judgehost && go build .

ENTRYPOINT ["./xoj_judgehost"]