PWD=`pwd`
GO_VERSION=1.13.4
TENSORFLOW_BRANCH=r2.1

all: go/tensorflow LICENSE

go/tensorflow: tensorflow go
	docker run --rm -v $(PWD)/tensorflow:/input/tensorflow:ro -v $(PWD)/go:/output/go/github.com/tensorflow/tensorflow golang:$(GO_VERSION)-alpine sh -c \
	'set -ex \
		&& apk add git protobuf-dev \
		&& cd /output/go \
		&& go get github.com/golang/protobuf/protoc-gen-go \
		&& protoc \
			-I/input/tensorflow \
			--go_out=plugins=grpc:/output/go \
			/input/tensorflow/tensorflow/core/framework/*.proto\
		&& protoc \
			-I/input/tensorflow \
			--go_out=plugins=paths=source_relative:/output/go \
			/input/tensorflow/tensorflow/core/example/*.proto \
		&& cd /output/go/github.com/tensorflow/tensorflow/tensorflow \
		&& go mod init github.com/tensorflow/tensorflow/tensorflow \
		&& go mod tidy'

LICENSE: tensorflow
	cp tensorflow/LICENSE LICENSE

tensorflow:
	git clone --branch $(TENSORFLOW_BRANCH) --depth 1 https://github.com/tensorflow/tensorflow.git tensorflow

go:
	mkdir go

clean:
	rm -rf go tensorflow LICENSE

.PHONY:	all clean
