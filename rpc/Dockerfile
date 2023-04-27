FROM golang:alpine as builder

RUN go install github.com/golang/protobuf/protoc-gen-go@latest
RUN export GO111MODULE=on && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN echo $GOPATH

RUN apk add --no-cache make protobuf-dev
RUN apk --no-cache add protobuf


WORKDIR /workspace

CMD ["protoc"]