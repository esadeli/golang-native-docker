FROM golang:latest

LABEL maintainer = "efrat.sadeli@gmail.com"

WORKDIR /

COPY . .

RUN go get -v ./...

RUN go build -o main .

EXPOSE 5051

CMD ["./main"]