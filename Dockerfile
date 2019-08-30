FROM golang:latest

LABEL maintainer = "efrat.sadeli@gmail.com"

WORKDIR /

COPY . .

EXPOSE 5051

CMD ["./main.go"]