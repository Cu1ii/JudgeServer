FROM gleaming/golang1.9.3:env
MAINTAINER cu1

ADD . ./src/Judge-server
WORKDIR ./src/Judge-server

ENV GO111MODULE on
ENV CGO_ENABLED 0

RUN mkdir build && cd build && cmake ../Judger && make && make install && cd ../xoj_judgehost && export GOPROXY=https://goproxy.cn && go mod tidy
# RUN go build xoj_judgehost
CMD ["go", "run", "main.go"]