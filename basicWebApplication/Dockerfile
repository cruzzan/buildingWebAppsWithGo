FROM golang

ADD . /go/src/app

WORKDIR /go/src/app

RUN go get -d -v ./...
RUN go install -v ./...
RUN go get github.com/astaxie/beego && go get github.com/beego/bee

CMD ["bee", "run"]

EXPOSE 8181
