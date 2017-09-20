FROM golang:1.9

ENV GOPATH /go/

RUN curl https://glide.sh/get | sh

RUN mkdir -p /go/src/github.com/bugcrowd/secrets
WORKDIR /go/src/github.com/bugcrowd/secrets
COPY . /go/src/github.com/bugcrowd/secrets

RUN make deps
RUN make build
CMD make test
