FROM golang:latest

RUN mkdir /go/src/rss-to-kindle
ADD . /go/src/rss-to-kindle
WORKDIR /go/src/rss-to-kindle

RUN go get github.com/fatih/color
RUN go get github.com/joho/godotenv
RUN go get github.com/mmcdole/gofeed
RUN go get github.com/nfnt/resize
RUN go get github.com/scorredoira/email

RUN go build -o main .
CMD ["/go/src/rss-to-kindle/main"]