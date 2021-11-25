FROM golang
MAINTAINER eicesoft
WORKDIR /go/src/
COPY . .
RUN go env -w GOPROXY="https://goproxy.io"
RUN go run .
EXPOSE 8889

ENTRYPOINT ["go", "run", "."]
