FROM golang:1.21.1-alpine3.17

ENV GOOS="linux"
ENV CGO_ENABLED=1
ENV GO111MODULE=on
ENV GOARCH=$TARGETARCH
ENV PACKAGES="ca-certificates git curl bash zsh make gcc g++ libc-dev librdkafka-dev pkgconf"
ENV ROOT /service-1

RUN apk update \
    && apk add --no-cache ${PACKAGES} \
    && update-ca-certificates

WORKDIR ${ROOT}

# RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
#     go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
#     go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
#     go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest && \
#     go install golang.org/x/tools/gopls@latest && \
#     go install honnef.co/go/tools/cmd/staticcheck@latest && \
#     go install github.com/cweill/gotests/gotests@latest && \
#     go install github.com/fatih/gomodifytags@latest && \

COPY ./ ./

WORKDIR ${ROOT}

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "-tags", "musl", "main.go"]
