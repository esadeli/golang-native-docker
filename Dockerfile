FROM golang:latest

LABEL maintainer = "efrat.sadeli@gmail.com"

WORKDIR /go/src/github.com/esadeli/golang-native-docker

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN /go/bin/dep ensure
RUN go build -o main .

EXPOSE 5051

CMD ["./main"]